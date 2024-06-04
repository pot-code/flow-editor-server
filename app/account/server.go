package account

import (
	account "flow-editor-server/gen/account"
	"flow-editor-server/gen/http/account/server"
	"flow-editor-server/internal/goa"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	zw "github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"goa.design/goa/v3/http"
)

type Server struct {
	s account.Service
	z *authorization.Authorizer[*oauth.IntrospectionContext]
}

// MountHttpServer implements goa.HttpServer.
func (s *Server) MountHttpServer(mux http.ResolverMuxer) {
	endpoints := account.NewEndpoints(s.s)
	srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, goa.ErrorFormatter)
	srv.Use(zw.New(s.z).RequireAuthorization())
	server.Mount(mux, srv)
}

var _ goa.HttpServer = (*Server)(nil)

func NewServer(s account.Service, z *authorization.Authorizer[*oauth.IntrospectionContext]) *Server {
	return &Server{s: s, z: z}
}
