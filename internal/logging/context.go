package logging

import (
	"context"
	"log"
)

type key string

const (
	logKey key = "logging"
)

// WithContext returns a new Context that carries a logger.
func WithContext(ctx context.Context, logPrefix string) (*log.Logger, context.Context) {
	logger := NewLogger(logPrefix)
	return logger, context.WithValue(ctx, logKey, logger)
}

// FromContext returns the logger and error channel values stored in ctx, if any.
func FromContext(ctx context.Context) (*log.Logger, bool) {
	logger, ok := ctx.Value(logKey).(*log.Logger)
	return logger, ok
}
