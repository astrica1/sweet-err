package sweeterr

import (
    "errors"
    "testing"
    "go.uber.org/zap/zaptest"
)

func TestSweetError(t *testing.T) {
    err := New(InternalError, "Database error", nil, errors.New("connection failed"))

    if err.Error() != "Code: 500, Message: Database error, Context: map[], Error: connection failed" {
        t.Fatalf("unexpected error format: %v", err.Error())
    }
}

func TestLogError(t *testing.T) {
    logger := zaptest.NewLogger(t)
    err := New(ValidationError, "Invalid input", nil, nil)
    LogError(logger, err)
}
