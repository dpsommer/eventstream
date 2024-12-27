package main

import (
	"context"

	"github.com/dpsommer/eventstream/internal/characters"
	"github.com/dpsommer/eventstream/internal/events"
	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/regions"
)

// TODO: define a proper config for the region map
// and load it in a nicer way
func populateMap(ctx context.Context) *regions.Map {
	m, ok := regions.FromContext(ctx)
	if !ok {
		m = regions.NewMap()
	}

	m.AddNode(regions.Town)
	m.AddNode(regions.Forest)

	m.AddEdge(regions.Town, regions.Forest, 3)

	return m
}

func eventLoop(ctx context.Context) {
	// TODO: kick off the event loop workers. how should this be defined?
	_, ctx = logging.WithContext(ctx, "event worker: ")

	character := characters.NewCharacter(ctx, "Duncan")

	for range 10 {
		events.Emit(ctx, &events.MovementEvent{
			Character:   character,
			Destination: regions.Forest,
		})
	}
}
