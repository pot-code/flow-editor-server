package account

import (
	account "flow-editor-server/gen/account"
	"flow-editor-server/gen/http/account/server"
	"flow-editor-server/internal/goa"

	"goa.design/goa/v3/http"
)

type Route struct {
	s account.Service
}

func (s *Route) MountRoute(mux http.ResolverMuxer) {
	endpoints := account.NewEndpoints(s.s)
	srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, goa.HttpErrorFormatter)
	server.Mount(mux, srv)
}

var _ goa.HttpRoute = (*Route)(nil)

func NewRoute(s account.Service) *Route {
	return &Route{s: s}
}
