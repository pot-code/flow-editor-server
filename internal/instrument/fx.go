package instrument

import "go.uber.org/fx"

var Module = fx.Module(
	"instrument",
	fx.Provide(
		NewOtelPropagator,
		NewOtelTracerProvider,
		NewOtelMeterProvider,
	),
)
