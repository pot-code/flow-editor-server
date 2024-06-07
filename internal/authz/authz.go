package authz

import (
	"context"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/rs/zerolog/log"
)

type Resource interface {
	ResourceID() string
	OwnerID() string
	Kind() string
}

type Authz[T Resource] struct {
	c *cerbos.GRPCClient
}

func NewAuthz[T Resource](c *cerbos.GRPCClient) *Authz[T] {
	return &Authz[T]{c: c}
}

func (az *Authz[T]) CheckPermission(ctx context.Context, obj T, action string) error {
	a := Context(ctx)
	p := cerbos.NewPrincipal(a.UserID, a.Roles...)
	r := cerbos.NewResource(obj.Kind(), obj.ResourceID())
	r.WithAttr("owner", obj.OwnerID())
	ok, err := az.c.IsAllowed(ctx, p, r, action)
	if err != nil {
		return err
	}
	log.Debug().
		Str("principal", a.UserID).
		Str("action", action).
		Str("obj", obj.ResourceID()).
		Str("kind", obj.Kind()).
		Bool("allowed", ok).
		Msg("authorization result")
	if !ok {
		return ErrUnauthorized
	}
	return nil
}
