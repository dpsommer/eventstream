package regions

import (
	"context"
)

type key string

const (
	mapKey key = "map"
)

// WithContext returns a new Context that carries a Map.
func WithContext(ctx context.Context) (*Map, context.Context) {
	regionMap := NewMap()
	return regionMap, context.WithValue(ctx, mapKey, regionMap)
}

// FromContext returns the regionMap and error channel values stored in ctx, if any.
func FromContext(ctx context.Context) (*Map, bool) {
	regionMap, ok := ctx.Value(mapKey).(*Map)
	return regionMap, ok
}
