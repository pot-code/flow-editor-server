package goa

import (
	"go.uber.org/fx"
)

var HttpModule = fx.Module(
	"goa",
	fx.Provide(
		NewHttpServer,
		NewMux,
	),
)
