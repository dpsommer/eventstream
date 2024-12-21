package events

import (
	"context"
	"time"

	"github.com/dpsommer/eventstream/internal/characters"
	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/regions"
)

type EventType int

type Event interface {
	Process(ctx context.Context)
}

type MovementEvent struct {
	Character   *characters.Character
	Destination regions.Location
}

func (e *MovementEvent) Process(ctx context.Context) {
	// TODO: nicer way to unpack context? into a single conf struct?
	m, ok := regions.FromContext(ctx)
	if !ok {
		m = regions.NewMap()
	}

	logger, ok := logging.FromContext(ctx)
	if !ok {
		logger = logging.NewLogger("movement event: ")
	}

	// need to lock here until we update the character location to avoid a
	// time-of-check to time-of-use race condition

	// TODO: this mutex is on the character, so we're blocking any other
	// safe-access behaviour until this event completes; it may be better to
	// define separate mutexes for specific synchronous behaviours
	e.Character.Lock()
	defer e.Character.Unlock()

	if e.Character.Location == e.Destination {
		// WARN
		logger.Printf("%s is already at %s\n", e.Character.Name, e.Destination)
		return
	}

	path, dist, err := m.Distance(e.Character.Location, e.Destination)
	if err != nil {
		// ERROR
		logger.Printf("failed to find distance between %s and %s: %s\n", e.Character.Location, e.Destination, err.Error())
		return
	}

	logger.Printf("%s is traveling from %s to %s along the following route:\n", e.Character.Name, e.Character.Location, e.Destination)
	// the list is in reverse order, so iterate backwards
	for i := len(path) - 1; i >= 0; i-- {
		logger.Printf("\t* %s\n", path[i].Value)
	}
	logger.Printf("The total distance is %d\n", dist)

	// TODO: should distance/time be linear? need to put some thought into
	// event durations mapping to real time as this is a bit limited/lazy
	duration := time.Duration(dist) * time.Second
	time.Sleep(duration)
	e.Character.Location = e.Destination
}
