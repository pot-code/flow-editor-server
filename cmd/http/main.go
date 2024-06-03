package main

import (
	"context"
	"flow-editor-server/app/account"
	"flow-editor-server/app/flow"
	_ "flow-editor-server/docs"
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
	"go.uber.org/fx"
	goahttp "goa.design/goa/v3/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to read config")
	}

	zitadel, err := authorization.New(
		context.Background(),
		zitadel.New(viper.GetString("ZITADEL_DOMAIN"), zitadel.WithInsecure(viper.GetString("ZITADEL_PORT"))),
		oauth.DefaultAuthorization("key.json"),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to zitadel")
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	if err := db.AutoMigrate(&flow.FlowModel{}, &account.AccountModel{}); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	addr := viper.GetString("HTTP_ADDR")
	fx.New(
		account.Module,

		fx.Provide(func(mux goahttp.ResolverMuxer, lc fx.Lifecycle) *http.Server {
			srv := &http.Server{
				Addr:    addr,
				Handler: mux,
			}
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					ln, err := net.Listen("tcp", addr)
					if err != nil {
						return err
					}
					log.Info().Str("addr", addr).Msg("HTTP server started")
					go srv.Serve(ln)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return srv.Shutdown(ctx)
				},
			})
			return srv
		}),
		fx.Provide(goahttp.NewMuxer),
		fx.Supply(db, zitadel),
		fx.Invoke(func(s *http.Server) {
		}),
	).Run()
}
