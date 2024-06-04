//go:build !goverter

package account

import (
	"flow-editor-server/gen/account"

	"go.uber.org/fx"
	"goa.design/goa/v3/http"
)

var Module = fx.Module(
	"account",
	fx.Provide(
		NewServer,
		fx.Annotate(NewService, fx.As(new(account.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
	fx.Invoke(func(s *Server, mux http.ResolverMuxer) {
		s.MountHttpServer(mux)
	}),
)
