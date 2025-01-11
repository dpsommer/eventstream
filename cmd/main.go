package main

import (
	"context"
	"time"

	"github.com/dpsommer/eventstream/internal/clock"
	"github.com/dpsommer/eventstream/internal/events"
)

func main() {
	// create a channel to wait on SIGINT/SIGTERM
	signals := setupSignals()

	// initialize global context
	// TODO: change this to a config struct
	ctx := populateMap(context.Background())
	_, ctx = events.WithContext(ctx)
	ctx = generateCharacters(ctx)

	// start the world clock
	go clock.StartWorldClock(ctx, &clock.WorldClockOptions{
		Tick:     100 * time.Millisecond,
		Shutdown: signals,
	})

	// block until we receive an interrupt or term
	<-signals
}
