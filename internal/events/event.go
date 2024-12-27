package events

import (
	"context"
	"fmt"
	"time"

	"github.com/dpsommer/eventstream/internal/characters"
	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/regions"
)

type EventType int

type Event interface {
	PreProcess(ctx context.Context)
	Process(ctx context.Context)
	Duration(ctx context.Context) (time.Duration, error)
	String() string
}

type MovementEvent struct {
	Character   *characters.Character
	Destination regions.Location
}

func (e *MovementEvent) PreProcess(ctx context.Context) {
	// XXX: this could also be used to update character state
	// (e.g. to "Moving") or trigger other events

	// need to lock here until we update the character location to avoid a
	// time-of-check to time-of-use race condition

	// TODO: this mutex is on the character, so we're blocking any other
	// safe-access behaviour until this event completes; it may be better to
	// define separate mutexes for specific synchronous behaviours
	e.Character.Lock()
}

func (e *MovementEvent) Process(ctx context.Context) {
	e.Character.Location = e.Destination
	// FIXME: acquiring the lock in preprocess and releaseing it here/in
	// Duration (or anywhere else with an error path) is very fragile and
	// highly error-prone. there must be a better way to handle bocking on
	// each actor
	e.Character.Unlock()
}

func (e *MovementEvent) Duration(ctx context.Context) (time.Duration, error) {
	m, ok := regions.FromContext(ctx)
	if !ok {
		m = regions.NewMap()
	}

	logger, ok := logging.FromContext(ctx)
	if !ok {
		logger = logging.NewLogger("movement event: ")
	}

	if e.Character.Location == e.Destination {
		err := fmt.Errorf("%s is already at %s", e.Character.Name, e.Destination)
		e.Character.Unlock()
		return time.Duration(0), err
	}

	path, dist, err := m.Distance(e.Character.Location, e.Destination)
	if err != nil {
		e.Character.Unlock()
		return time.Duration(0), err
	}

	logger.Printf("%s is traveling from %s to %s along the following route:\n", e.Character.Name, e.Character.Location, e.Destination)
	// the path is in reverse order, so iterate backwards
	for i := len(path) - 1; i >= 0; i-- {
		logger.Printf("\t* %s\n", path[i].Value)
	}
	logger.Printf("The total distance is %d\n", dist)

	// TODO: should distance/time be linear? need to put some thought into
	// event durations mapping to real time as this is a bit limited/lazy
	return time.Duration(dist) * time.Second, nil
}

func (e *MovementEvent) String() string {
	return fmt.Sprintf("move character %s to %s", e.Character.Name, e.Destination)
}
