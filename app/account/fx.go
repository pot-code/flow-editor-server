//go:build !goverter

package account

import (
	"flow-editor-server/gen/account"
	"flow-editor-server/internal/goa"

	"go.uber.org/fx"
)

var HttpModule = fx.Module(
	"account",
	fx.Provide(
		fx.Annotate(NewRoute, fx.As(new(goa.HttpRoute)), fx.ResultTags(`group:"routes"`)),
		fx.Annotate(NewService, fx.As(new(account.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
)
