package sweeterr

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func TraceError(span trace.Span, err error) {
	if sweetErr, ok := err.(*SweetError); ok {
		span.SetAttributes(attribute.Int("error.code", int(sweetErr.Code)))
		span.SetAttributes(attribute.String("error.message", sweetErr.Message))
		for key, value := range sweetErr.Context {
			span.SetAttributes(attribute.String(fmt.Sprintf("error.context.%s", key), fmt.Sprintf("%v", value)))
		}
	}
	span.RecordError(err)
}
