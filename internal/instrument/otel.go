package instrument

import (
	"context"
	"flow-editor-server/internal/config"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
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
