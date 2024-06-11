//go:build !goverter

package account

import (
	"context"
	"flow-editor-server/gen/account"
	"flow-editor-server/gen/http/account/server"
	"flow-editor-server/internal/goa"

	"go.uber.org/fx"
	"goa.design/goa/v3/http"
	"gorm.io/gorm"
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
	fx.Invoke(func(as account.Service, muxer http.ResolverMuxer) {
		endpoints := account.NewEndpoints(as)
		srv := server.New(endpoints, muxer, http.RequestDecoder, http.ResponseEncoder, nil, goa.HttpErrorFormatter)
		server.Mount(muxer, srv)
	}),
	fx.Invoke(func(db *gorm.DB, l fx.Lifecycle) {
		l.Append(fx.Hook{
			OnStart: func(ctx context.Context) error {
				if err := db.AutoMigrate(&Account{}); err != nil {
					return err
				}
				if err := db.Where("name=?", "admin").Attrs(&Role{Name: "admin"}).FirstOrCreate(&Role{}).Error; err != nil {
					return err
				}
				if err := db.Where("name=?", "user").Attrs(&Role{Name: "user"}).FirstOrCreate(&Role{}).Error; err != nil {
					return err
				}
				return nil
			},
		})
	}),
)
