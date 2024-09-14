package sweeterr

import (
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

func ErrorHandlerMiddleware(logger Logger, tracer trace.Tracer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					err, ok := rec.(error)
					if !ok {
						err = fmt.Errorf("panic: %v", rec)
					}
					span := trace.SpanFromContext(r.Context())
					TraceError(span, err)
					LogError(logger, err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
