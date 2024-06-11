package instrument

import (
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

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
