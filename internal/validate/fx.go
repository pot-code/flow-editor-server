package validate

import "go.uber.org/fx"

var Module = fx.Module(
	"validate",
	fx.Provide(
		NewValidator,
		NewTranslator,
	),
)
