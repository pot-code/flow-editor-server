package instrument

import (
	"context"
	"flow-editor-server/internal/config"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewOtelPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
}

func NewOtelTracerProvider(config *config.HttpConfig) (*trace.TracerProvider, error) {
	e, err := otlptracegrpc.New(context.Background(), otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(config.OtelCollector))
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer provider: %w", err)
	}
	tp := trace.NewTracerProvider(trace.WithBatcher(e))
	return tp, nil
}

func NewOtelMeterProvider(config *config.HttpConfig) (*metric.MeterProvider, error) {
	e, err := otlpmetricgrpc.New(context.Background(), otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint(config.OtelCollector))
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics exporter: %w", err)
	}
	mp := metric.NewMeterProvider(metric.WithReader(metric.NewPeriodicReader(e)))
	return mp, nil
}
