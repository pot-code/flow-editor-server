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

type Server struct {
	s flow.Service
	v *validator.Validate
	t ut.Translator
	z *authorization.Authorizer[*oauth.IntrospectionContext]
}

// MountHttpServer implements goa.Server.
func (s *Server) MountHttpServer(mux http.ResolverMuxer) {
	endpoints := flow.NewEndpoints(s.s)
	endpoints.Use(goa.ValidatePayload(s.v, s.t))
	srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, goa.ErrorFormatter)
	srv.Use(zw.New(s.z).RequireAuthorization())
	server.Mount(mux, srv)
}

var _ goa.HttpServer = (*Server)(nil)

func NewServer(
	s flow.Service,
	v *validator.Validate,
	t ut.Translator,
	z *authorization.Authorizer[*oauth.IntrospectionContext],
) *Server {
	return &Server{s: s, v: v, t: t, z: z}
}
