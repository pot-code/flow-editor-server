package main

import (
	"context"
	"flow-editor-server/app/account"
	"flow-editor-server/app/flow"
	"flow-editor-server/internal/config"
	"flow-editor-server/internal/validate"
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
	config := config.NewHttpConfig()

	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	if err := db.AutoMigrate(&flow.FlowModel{}, &account.AccountModel{}); err != nil {
		log.Fatal().Err(err).Msg("failed to migrate database")
	}

	zitadel, err := authorization.New(
		context.Background(),
		zitadel.New(viper.GetString("ZITADEL_DOMAIN"), zitadel.WithInsecure(viper.GetString("ZITADEL_PORT"))),
		oauth.DefaultAuthorization("key.json"),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to zitadel")
	}

	fx.New(
		account.Module,
		flow.Module,

		fx.Supply(db, zitadel, config),
		fx.Provide(goahttp.NewMuxer, validate.NewTranslator, validate.NewValidator, newHttpServer),
		fx.Invoke(func(s *http.Server) {}),
	).Run()
}

func newHttpServer(mux goahttp.ResolverMuxer, config *config.HttpConfig, lc fx.Lifecycle) *http.Server {
	srv := &http.Server{
		Addr:    config.Addr,
		Handler: mux,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", config.Addr)
			if err != nil {
				return err
			}
			log.Info().Str("addr", config.Addr).Msg("HTTP server started")
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
