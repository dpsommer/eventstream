package characters

import (
	"context"
)

type key string

const (
	obsKey key = "observer"
)

// WithContext returns a new Context that carries an observer.
func WithContext(ctx context.Context) (*Observer, context.Context) {
	observer := NewObserver(ctx)
	return observer, context.WithValue(ctx, obsKey, observer)
}

// FromContext returns the observer and error channel values stored in ctx, if any.
func FromContext(ctx context.Context) (*Observer, bool) {
	observer, ok := ctx.Value(obsKey).(*Observer)
	return observer, ok
}
