package main

import (
	"context"

	"github.com/dpsommer/eventstream/internal/events"
	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/regions"
)

func main() {
	logger := logging.NewLogger("main: ")

	// create a channel to wait on SIGINT/SIGTERM
	signals := setupSignals()

	errs := make(chan error)
	defer close(errs)

	// XXX: AddListener creates background workers
	// this side effect is convenient, but might be better
	// broken out into separate functions e.g. if we need
	// references to the listeners
	logger.Printf("starting %d event listeners\n", events.MaxWorkers)

	_, ctx := regions.WithContext(context.Background())
	_, ctx = logging.WithContext(ctx, "event worker: ")
	populateMap(ctx)

	// start the event loop
	go eventLoop(ctx)

	// block until we receive an interrupt or term
	<-signals
}
