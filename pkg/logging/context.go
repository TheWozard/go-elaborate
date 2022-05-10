package logging

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type loogerContextKey int

const (
	TimestampField  = "timestamp"
	TimestampFormat = time.RFC3339Nano

	loggerKey loogerContextKey = iota
)

var logger *zap.Logger

func init() {
	logger, _ = zap.NewProduction()
}

// NewContext wraps the logger in the context with new fields
func NewContext(ctx context.Context, fields ...zap.Field) context.Context {
	return context.WithValue(ctx, loggerKey, From(ctx).With(fields...))
}

// From attempts to pull the logger off the context, returns a default logger in the event of failure
func From(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}
