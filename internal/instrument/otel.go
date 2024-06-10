package instrument

import (
	"context"
	"flow-editor-server/internal/config"
	"fmt"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

func NewOtelPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
}

func NewOtelTracerExporter(config *config.HttpConfig) (*otlptrace.Exporter, error) {
	e, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(config.OtelCollector),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer provider: %w", err)
	}
	return e, nil
}

func NewOtelMeterExporter(config *config.HttpConfig) (*otlpmetricgrpc.Exporter, error) {
	e, err := otlpmetricgrpc.New(context.Background(), otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint(config.OtelCollector))
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}
	return e, nil
}

type ZerologHook struct{}

func (h *ZerologHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	ctx := e.GetCtx()
	span := trace.SpanContextFromContext(ctx)
	if span.HasTraceID() {
		spanId := span.TraceID().String()
		e.Str("trace_id", spanId)
	}
}

func NewZerologHook() zerolog.Hook {
	return &ZerologHook{}
}
