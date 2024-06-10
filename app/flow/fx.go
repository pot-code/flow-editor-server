//go:build !goverter

package flow

import (
	"context"
	"flow-editor-server/gen/flow"
	"flow-editor-server/internal/goa"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"gorm.io/gorm"
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
		fx.Annotate(otel.Tracer("flow"), fx.As(new(trace.Tracer))),
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
	fx.Invoke(func(db *gorm.DB, l fx.Lifecycle) {
		l.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				return db.AutoMigrate(
					&Flow{},
				)
			},
		})
	}),
)
