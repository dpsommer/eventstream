package main

import (
	"context"
	"time"

	"github.com/dpsommer/eventstream/internal/events"
	"github.com/dpsommer/eventstream/internal/regions"
)

func main() {
	// create a channel to wait on SIGINT/SIGTERM
	signals := setupSignals()

	errs := make(chan error)
	defer close(errs)

	// initialize global context
	// TODO: change this to a config struct
	_, ctx := regions.WithContext(context.Background())
	sched, ctx := events.WithContext(ctx)
	populateMap(ctx)

	// start the world clock
	go events.StartWorldClock(ctx, &events.WorldClockOptions{
		Tick:      100 * time.Millisecond,
		Shutdown:  signals,
		Scheduler: sched,
	})

	// start the event loop
	go eventLoop(ctx)

	// block until we receive an interrupt or term
	<-signals
}
