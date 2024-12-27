package events

import (
	"context"
)

type key string

const (
	schedKey key = "scheduler"
)

// WithContext returns a new Context that carries a scheduler.
func WithContext(ctx context.Context) (*Scheduler, context.Context) {
	scheduler := NewScheduler(ctx)
	return scheduler, context.WithValue(ctx, schedKey, scheduler)
}

// FromContext returns the scheduler and error channel values stored in ctx, if any.
func FromContext(ctx context.Context) (*Scheduler, bool) {
	scheduler, ok := ctx.Value(schedKey).(*Scheduler)
	return scheduler, ok
}
