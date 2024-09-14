package sweeterr

import (
	"fmt"
)

type ErrorCode int

type SweetError struct {
	Code    ErrorCode
	Message string
	Context map[string]interface{}
	Err     error
}

const (
	ValidationError ErrorCode = 400
	InternalError   ErrorCode = 500
	NotFoundError   ErrorCode = 404
)

func New(code ErrorCode, message string, context map[string]interface{}, err error) *SweetError {
	return &SweetError{
		Code:    code,
		Message: message,
		Context: context,
		Err:     err,
	}
}

func (e *SweetError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Code: %d, Message: %s, Context: %v, Error: %v", e.Code, e.Message, e.Context, e.Err)
	}
	return fmt.Sprintf("Code: %d, Message: %s, Context: %v", e.Code, e.Message, e.Context)
}

func (e *SweetError) Unwrap() error {
	return e.Err
}
