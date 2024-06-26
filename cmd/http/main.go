package main

import (
	"context"
	"flow-editor-server/app"
	"flow-editor-server/internal/authn"
	"flow-editor-server/internal/authz"
	"flow-editor-server/internal/config"
	"flow-editor-server/internal/orm"
	"flow-editor-server/internal/validate"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
	ghttp "goa.design/goa/v3/http"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fx.New(
		validate.Module, app.Module,

		fx.Provide(config.NewHttpConfig, authz.NewCerbosClient, orm.NewGormDB),

		// http muxer
		fx.Provide(fx.Annotate(func(
			config *config.HttpConfig,
			l fx.Lifecycle,
		) ghttp.ResolverMuxer {
			muxer := ghttp.NewMuxer()
			muxer.Use(alice.New(
				hlog.NewHandler(log.Logger),
				hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
					hlog.FromRequest(r).Debug().
						Ctx(r.Context()).
						Str("method", r.Method).
						Str("path", r.URL.Path).
						Int("status", status).
						Int("size", size).
						Dur("duration", duration).
						Send()
				}),
				authn.JwtValidation(config.OidcJwkProvider, config.OidcProvider, config.OidcApiID),
			).Then)
			return muxer
		})),

		// http server
		fx.Invoke(func(mux ghttp.ResolverMuxer, config *config.HttpConfig, l fx.Lifecycle) {
			srv := &http.Server{Addr: config.Addr, Handler: mux}
			if config.Debug {
				zerolog.SetGlobalLevel(zerolog.DebugLevel)
			}
			l.Append(fx.Hook{
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
					return srv.Close()
				},
			})
		}),

		fx.StartTimeout(30*time.Second),
	).Run()
}
