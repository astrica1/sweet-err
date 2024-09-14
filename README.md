# SweetError

Golang Error Management Library

**SweetError** is a powerful and extensible error handling library for Golang, designed to provide a structured way to manage errors, logging, and tracing. It is built to support complex applications where effective error handling is crucial for maintaining stability and observability.

## Features

- **Structured Error Handling**: Provides a customizable error structure with support for error codes, messages, context, and wrapped errors.
- **Logging Integration**: Easily integrates with popular logging libraries like `zap` to log detailed error information.
- **Tracing Integration**: Supports tracing using OpenTelemetry, allowing you to track errors across distributed systems.
- **HTTP Middleware**: Middleware for handling errors in HTTP applications, logging, and tracing in one go.
- **Customizable**: Designed to be easily extended and adapted for various error-handling use cases.

## Installation

To install the library, simply use `go get`:

```bash
go get github.com/astrica1/sweet-error
```

## Getting Started

### 1. Creating and Using Errors

You can create structured errors using the New function and wrap existing errors with additional context:

```go
import "github.com/astrica1/sweet-error"

err := sweeterr.New(sweeterr.InternalError, "Database error", map[string]interface{}{
    "query": "SELECT * FROM users",
}, originalError)

fmt.Println(err.Error())

```

### 2. Logging Errors

SweetError integrates with zap or any other logger that implements the Logger interface. Here’s an example of logging an error:

```go
import (
    "go.uber.org/zap"
    "github.com/astrica1/sweet-error"
)

logger, _ := zap.NewProduction()
err := sweeterr.New(sweeterr.ValidationError, "Invalid input", nil, nil)
sweeterr.LogError(logger, err)

```

### 3. Tracing Errors

For distributed tracing, you can record error information into OpenTelemetry spans:

```go
import (
    "go.opentelemetry.io/otel"
    "github.com/astrica1/sweet-error"
)

tracer := otel.Tracer("example-tracer")
ctx, span := tracer.Start(context.Background(), "example-operation")
defer span.End()

err := sweeterr.New(sweeterr.NotFoundError, "User not found", nil, nil)
sweeterr.TraceError(span, err)

```

### 4. Middleware for HTTP Servers

The library provides a middleware for handling errors in HTTP applications. It logs and traces errors automatically.

```go
import (
    "net/http"
    "go.uber.org/zap"
    "go.opentelemetry.io/otel"
    "github.com/astrica1/sweet-error"
)

logger, _ := zap.NewProduction()
tracer := otel.Tracer("example-tracer")

middleware := sweeterr.ErrorHandlerMiddleware(logger, tracer)

http.Handle("/", middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    panic("unexpected error")
})))

http.ListenAndServe(":8080", nil)

```

## Customization

You can extend the library to suit your needs, such as creating custom error types, adding more sophisticated logging, or handling specific error cases.

### Example: Custom Error Type

```go
type DatabaseError struct {
    errorlib.CustomError
}

func NewDatabaseError(query string, err error) *DatabaseError {
    return &DatabaseError{
        CustomError: sweeterr.SweetError{
            Code:    sweeterr.InternalError,
            Message: "Database error occurred",
            Context: map[string]interface{}{
                "query": query,
            },
            Err:     err,
        },
    }
}


```

## Running Tests

To run the tests for this library, you can use the standard Go test command:

```bash
go test ./...
```

## Contributing

Contributions are welcome! If you’d like to contribute to the project, feel free to submit a pull request or open an issue.

### Steps to Contribute:

1. Fork the repository
2. Create a new feature branch
3. Commit your changes
4. Push the branch and create a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Contact

If you have any questions or suggestions, feel free to open an issue or contact the maintainers.
