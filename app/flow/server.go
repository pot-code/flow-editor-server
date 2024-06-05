package flow

import (
	"flow-editor-server/gen/flow"
	"flow-editor-server/gen/http/flow/server"
	"flow-editor-server/internal/goa"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	zw "github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"goa.design/goa/v3/http"
)

type Route struct {
	s flow.Service
	v *validator.Validate
	t ut.Translator
	z *authorization.Authorizer[*oauth.IntrospectionContext]
}

// MountRoute implements goa.Server.
func (s *Route) MountRoute(mux http.ResolverMuxer) {
	endpoints := flow.NewEndpoints(s.s)
	endpoints.Use(goa.ValidatePayload(s.v, s.t))
	srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, goa.ErrorFormatter)
	srv.Use(zw.New(s.z).RequireAuthorization())
	server.Mount(mux, srv)
}

var _ goa.HttpRoute = (*Route)(nil)

func NewRoute(
	s flow.Service,
	v *validator.Validate,
	t ut.Translator,
	z *authorization.Authorizer[*oauth.IntrospectionContext],
) *Route {
	return &Route{s: s, v: v, t: t, z: z}
}
