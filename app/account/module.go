//go:build !goverter

package account

import (
	"flow-editor-server/gen/account"
	"flow-editor-server/gen/http/account/server"

	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"go.uber.org/fx"
	"goa.design/goa/v3/http"
)

var Module = fx.Module(
	"account",
	fx.Provide(
		fx.Annotate(NewService, fx.As(new(account.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
	fx.Invoke(func(s account.Service, mux http.ResolverMuxer, zitadel *authorization.Authorizer[*oauth.IntrospectionContext]) {
		endpoints := account.NewEndpoints(s)
		srv := server.New(endpoints, mux, http.RequestDecoder, http.ResponseEncoder, nil, nil)
		srv.Use(middleware.New(zitadel).RequireAuthorization())
		server.Mount(mux, srv)
	}),
)
