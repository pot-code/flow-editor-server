//go:build !goverter

package flow

import (
	"context"
	aa "flow-editor-server/app/account"
	"flow-editor-server/gen/account"
	"flow-editor-server/gen/flow"
	"flow-editor-server/gen/http/flow/server"
	"flow-editor-server/internal/goa"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"goa.design/goa/v3/http"
	"gorm.io/gorm"
)

var HttpModule = fx.Module(
	"flow",
	fx.Provide(
		NewAuthz,
		fx.Annotate(NewService, fx.As(new(flow.Service))),
	),
	fx.Supply(
		fx.Private,
		fx.Annotate(new(ConverterImpl), fx.As(new(Converter))),
	),
	fx.Invoke(func(s flow.Service, as account.Service, v *validator.Validate, t ut.Translator, muxer http.ResolverMuxer) {
		endpoints := flow.NewEndpoints(s)
		endpoints.Use(goa.ValidatePayload(v, t))
		srv := server.New(endpoints, muxer, http.RequestDecoder, http.ResponseEncoder, nil, goa.HttpErrorFormatter)
		srv.Use(aa.Middleware(as))
		server.Mount(muxer, srv)
	}),
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
