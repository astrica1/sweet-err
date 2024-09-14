package sweeterr

import (
	"go.uber.org/zap"
)

type Logger interface {
	Error(msg string, fields ...zap.Field)
}

func LogError(logger Logger, err error) {
	if sweetErr, ok := err.(*SweetError); ok {
		logger.Error("Sweet Error",
			zap.Int("Code", int(sweetErr.Code)),
			zap.String("Message", sweetErr.Message),
			zap.Any("Context", sweetErr.Context),
			zap.Error(sweetErr.Err),
		)
	} else {
		logger.Error(err.Error())
	}
}
