package main

import (
	"context"
	"flow-editor-server/app/account"
	"flow-editor-server/app/flow"
	"flow-editor-server/internal/authn"
	"flow-editor-server/internal/authz"
	"flow-editor-server/internal/config"
	"flow-editor-server/internal/goa"
	"flow-editor-server/internal/orm"
	"flow-editor-server/internal/validate"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fx.New(
		goa.HttpModule,
		account.HttpModule,
		flow.HttpModule,

		fx.Provide(
			validate.NewTranslator,
			validate.NewValidator,
			config.NewHttpConfig,
			authz.NewCerbosClient,
			authn.NewZitadelClient,
			orm.NewGormDB,
		),
		// db migrations
		fx.Invoke(func(db *gorm.DB, l fx.Lifecycle) {
			l.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					if err := db.AutoMigrate(
						&flow.Flow{},
						&account.Account{},
					); err != nil {
						return fmt.Errorf("failed to migrate database: %w", err)
					}
					return nil
				},
			})
		}),
		// http server
		fx.Invoke(func(*http.Server) {}),
		// app logging
		// should be invoked at the end
		// or the diagnose logs will be mixed
		fx.Invoke(func(l fx.Lifecycle) {
			output, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			l.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					return output.Close()
				},
				OnStart: func(ctx context.Context) error {
					if err != nil {
						return fmt.Errorf("failed to open app log file: %w", err)
					}
					log.Logger = log.Output(output)
					return nil
				},
			})
		}),
		fx.StartTimeout(30*time.Second),
	).Run()
}
