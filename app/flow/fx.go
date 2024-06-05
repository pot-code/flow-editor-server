//go:build !goverter

package flow

import (
	"flow-editor-server/gen/flow"

	"go.uber.org/fx"
	"goa.design/goa/v3/http"
)

var Module = fx.Module(
	"flow",
	fx.Provide(
		NewRoute,
		fx.Annotate(NewService, fx.As(new(flow.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
	fx.Invoke(func(s *Route, mux http.ResolverMuxer) {
		s.MountRoute(mux)
	}),
)
