//go:build !goverter

package flow

import (
	"flow-editor-server/gen/flow"
	"flow-editor-server/internal/goa"

	"go.uber.org/fx"
)

var HttpModule = fx.Module(
	"flow",
	fx.Provide(
		NewAuthz,
		fx.Annotate(NewRoute, fx.As(new(goa.HttpRoute)), fx.ResultTags(`group:"routes"`)),
		fx.Annotate(NewService, fx.As(new(flow.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
)
