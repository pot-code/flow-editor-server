package flow

import (
	"context"
	"errors"
	"flow-editor-server/internal/authz"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type Authz struct {
	a  *authz.Authz[*Flow]
	c  *cerbos.GRPCClient
	db *gorm.DB
}

func NewAuthz(c *cerbos.GRPCClient, db *gorm.DB) *Authz {
	return &Authz{a: authz.NewAuthz[*Flow](c), c: c, db: db}
}

func (a *Authz) CheckPermission(ctx context.Context, obj *Flow, action string) error {
	return a.a.CheckPermission(ctx, obj, action)
}

func (a *Authz) CheckCreatePermission(ctx context.Context) error {
	ac := authz.Context(ctx)
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
