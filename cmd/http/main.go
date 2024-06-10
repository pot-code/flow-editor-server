package main

import (
	"context"
	"flow-editor-server/app/account"
	"flow-editor-server/app/flow"
	"flow-editor-server/internal/config"
	"flow-editor-server/internal/validate"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/cerbos/cerbos-sdk-go/cerbos"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	zw "github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
	"go.uber.org/fx"
	ghttp "goa.design/goa/v3/http"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fx.New(
		account.Module,
		flow.Module,

		fx.Provide(
			validate.NewTranslator,
			validate.NewValidator,
			config.NewHttpConfig,
			newHttpServer,
			newGormDB,
			newMux,
			newZitadelClient,
			newCerbosClient,
		),
		fx.Invoke(func(s *http.Server, l fx.Lifecycle) {
			output, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			l.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					if err != nil {
						return fmt.Errorf("failed to open app log file: %w", err)
					}
					log.Logger = log.Output(output)
					return nil
				},
				OnStop: func(ctx context.Context) error {
					return output.Close()
				},
			})
		}),
		fx.StartTimeout(30*time.Second),
	).Run()
}

func newGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database")
	}
	if err := db.AutoMigrate(
		&flow.Flow{},
		&account.Account{},
	); err != nil {
		return nil, fmt.Errorf("failed to migrate database")
	}
	return db, nil
}

func newZitadelClient(config *config.HttpConfig) (*authorization.Authorizer[*oauth.IntrospectionContext], error) {
	z, err := authorization.New(
		context.Background(),
		zitadel.New(config.ZitadelDomain, zitadel.WithInsecure(config.ZitadelPort)),
		oauth.DefaultAuthorization("key.json"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to zitadel: %w", err)
	}
	return z, err
}

func newMux(z *authorization.Authorizer[*oauth.IntrospectionContext], config *config.HttpConfig, l fx.Lifecycle) (ghttp.ResolverMuxer, error) {
	output, err := os.OpenFile("logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open access log file: %w", err)
	}
	mux := ghttp.NewMuxer()
	mux.Use(alice.New(
		hlog.NewHandler(zerolog.New(output).With().Timestamp().Logger()),
		hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
			hlog.FromRequest(r).Info().
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Int("status", status).
				Int("size", size).
				Dur("duration", duration).
				Msg("")
		}),
		zw.New(z).RequireAuthorization(),
	).Then)

	l.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return output.Close()
		},
	})
	return mux, nil
}

func newHttpServer(mux ghttp.ResolverMuxer, config *config.HttpConfig, lc fx.Lifecycle) *http.Server {
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

func newCerbosClient(config *config.HttpConfig) (*cerbos.GRPCClient, error) {
	c, err := cerbos.New(config.CerobsAddr, cerbos.WithPlaintext())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to cerbos: %w", err)
	}
	return c, err
}
