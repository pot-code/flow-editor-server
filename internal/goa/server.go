package goa

import (
	"context"
	"flow-editor-server/internal/config"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization"
	"github.com/zitadel/zitadel-go/v3/pkg/authorization/oauth"
	zw "github.com/zitadel/zitadel-go/v3/pkg/http/middleware"
	"go.uber.org/fx"
	ghttp "goa.design/goa/v3/http"
)

func NewMux(z *authorization.Authorizer[*oauth.IntrospectionContext], l fx.Lifecycle) (ghttp.ResolverMuxer, error) {
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

func NewHttpServer(mux ghttp.ResolverMuxer, config *config.HttpConfig, l fx.Lifecycle) *http.Server {
	srv := &http.Server{
		Addr:    config.Addr,
		Handler: mux,
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
	return srv
}
