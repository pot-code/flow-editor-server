package main

import (
	"context"
	"flow-editor-server/app"
	"flow-editor-server/internal/authn"
	"flow-editor-server/internal/authz"
	"flow-editor-server/internal/config"
	"flow-editor-server/internal/goa"
	"flow-editor-server/internal/instrument"
	"flow-editor-server/internal/orm"
	"flow-editor-server/internal/validate"
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
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.20.0"
	"go.uber.org/fx"
	ghttp "goa.design/goa/v3/http"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger.Level(zerolog.DebugLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	fx.New(
		instrument.Module,
		validate.Module,
		app.Module,

		fx.Provide(
			config.NewHttpConfig,
			authz.NewCerbosClient,
			authn.NewZitadelClient,
			orm.NewGormDB,
		),

		// trace provider
		fx.Provide(func(e *otlptrace.Exporter) *trace.TracerProvider {
			res := resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("flow-editor"),
				semconv.ServiceVersionKey.String("1.0.0"),
			)
			tp := trace.NewTracerProvider(trace.WithBatcher(e), trace.WithResource(res))
			return tp
		}),

		// http muxer
		fx.Provide(fx.Annotate(func(
			routes []goa.HttpRoute,
			tp *trace.TracerProvider,
			z *authorization.Authorizer[*oauth.IntrospectionContext],
			l fx.Lifecycle,
		) (ghttp.ResolverMuxer, error) {
			muxer := ghttp.NewMuxer()
			output, err := os.OpenFile("logs/access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				return nil, fmt.Errorf("failed to open log file: %w", err)
			}
			muxer.Use(alice.New(
				otelhttp.NewMiddleware("http.request",
					otelhttp.WithTracerProvider(tp),
					otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
						return fmt.Sprintf("%s %s", r.Method, r.URL.Path)
					}),
				),
				hlog.NewHandler(log.Output(output).Hook(instrument.NewZerologHook())),
				hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
					hlog.FromRequest(r).Info().
						Ctx(r.Context()).
						Str("method", r.Method).
						Str("path", r.URL.Path).
						Int("status", status).
						Int("size", size).
						Dur("duration", duration).
						Msg("")
				}),
				zw.New(z).RequireAuthorization(),
			).Then)

			for _, route := range routes {
				route.MountRoute(muxer)
			}

			l.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					return output.Close()
				},
			})
			return muxer, nil
		}, fx.ParamTags(`group:"routes"`))),

		// opentelemetry
		fx.Invoke(func(p propagation.TextMapPropagator, tp *trace.TracerProvider, l fx.Lifecycle) {
			otel.SetTextMapPropagator(p)
			otel.SetTracerProvider(tp)

			l.Append(fx.Hook{
				OnStop: func(ctx context.Context) error {
					return tp.Shutdown(ctx)
				},
			})
		}),

		// http server
		fx.Invoke(func(mux ghttp.ResolverMuxer, config *config.HttpConfig, l fx.Lifecycle) {
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
		}),

		// app logging
		// should be invoked after other components
		// or the diagnose logs will be mixed
		fx.Invoke(func(l fx.Lifecycle) {
			output, err := os.OpenFile("logs/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			l.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					if err != nil {
						return fmt.Errorf("failed to open app log file: %w", err)
					}
					log.Logger = log.Output(output).Hook(instrument.NewZerologHook())
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
