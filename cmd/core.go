package main

import (
	"context"

	"github.com/dpsommer/eventstream/internal/characters"
	"github.com/dpsommer/eventstream/internal/regions"
)

// TODO: define a proper config for the region map
// and load it in a nicer way
func populateMap(ctx context.Context) context.Context {
	m, ok := regions.FromContext(ctx)
	if !ok {
		m, ctx = regions.WithContext(ctx)
	}

	m.AddNode(regions.Town)
	m.AddNode(regions.Forest)

	m.AddEdge(regions.Town, regions.Forest, 3)

	return ctx
}

func generateCharacters(ctx context.Context) context.Context {
	o, ok := characters.FromContext(ctx)
	if !ok {
		o, ctx = characters.WithContext(ctx)
	}

	for range 10 {
		o.AddCharacter(regions.Town)
	}

	return ctx
}
