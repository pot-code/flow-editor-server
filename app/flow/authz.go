package flow

import (
	"context"
	"errors"
	"flow-editor-server/app/account"
	"flow-editor-server/internal/authz"
	"strconv"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Authz struct {
	c  *cerbos.GRPCClient
	db *gorm.DB
}

func NewAuthz(c *cerbos.GRPCClient, db *gorm.DB) *Authz {
	return &Authz{c: c, db: db}
}

func (az *Authz) CheckPermission(ctx context.Context, obj *Flow, action string) error {
	a := account.FromContext(ctx)
	ri := strconv.Itoa(int(obj.ID))
	p := cerbos.NewPrincipal(a.UserID, a.Roles...)
	r := cerbos.NewResource("flow", ri)
	r.WithAttr("owner", obj.Owner)
	ok, err := az.c.IsAllowed(ctx, p, r, action)
	if err != nil {
		return err
	}
	log.Debug().
		Ctx(ctx).
		Str("principal", a.UserID).
		Str("action", action).
		Str("obj", ri).
		Str("kind", "flow").
		Bool("allowed", ok).
		Msg("authorization result")
	if !ok {
		return authz.NewUnAuthorizedError(errors.New("未授权的操作"))
	}
	return nil
}

func (a *Authz) CheckCreatePermission(ctx context.Context) error {
	ac := account.FromContext(ctx)
	P := cerbos.NewPrincipal(ac.UserID, ac.Roles...)
	P.WithAttr("membership", ac.Membership)

	R := cerbos.NewResource("flow", "*")
	var total int64
	if err := a.db.Model(&Flow{}).Where("owner = ?", ac.UserID).Count(&total).Error; err != nil {
		return err
	}
	R.WithAttr("total", total)
	ok, err := a.c.IsAllowed(ctx, P, R, "create")
	if err != nil {
		return err
	}
	log.Debug().
		Ctx(ctx).
		Str("principal", ac.UserID).
		Int("membership", ac.Membership).
		Int64("total", total).
		Str("action", "create").
		Bool("allowed", ok).
		Msg("authorization result")
	if !ok {
		return authz.NewUnAuthorizedError(errors.New("flow 创建数已达到上限"))
	}
	return nil
}
