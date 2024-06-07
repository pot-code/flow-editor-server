package flow

import (
	"flow-editor-server/gen/account"
	"flow-editor-server/gen/flow"
	"flow-editor-server/gen/http/flow/server"
	"flow-editor-server/internal/authz"
	"flow-editor-server/internal/goa"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"goa.design/goa/v3/http"
)

type Route struct {
	s  flow.Service
	v  *validator.Validate
	t  ut.Translator
	as account.Service
}

func (s *Route) MountRoute(mux http.ResolverMuxer) {
	endpoints := flow.NewEndpoints(s.s)
	endpoints.Use(goa.ValidatePayload(s.v, s.t))
	srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, goa.ErrorFormatter)
	srv.Use(authz.Middleware(s.as))
	server.Mount(mux, srv)
}

var _ goa.HttpRoute = (*Route)(nil)

func NewRoute(
	s flow.Service,
	as account.Service,
	v *validator.Validate,
	t ut.Translator,
) *Route {
	return &Route{s: s, v: v, t: t, as: as}
}
