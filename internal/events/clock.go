package events

import (
	"context"
	"os"
	"time"
)

type WorldClockOptions struct {
	Tick      time.Duration
	Shutdown  <-chan os.Signal
	Scheduler *Scheduler
}

func StartWorldClock(ctx context.Context, opts *WorldClockOptions) {
	ticker := time.Tick(opts.Tick)

	for {
		select {
		case <-opts.Shutdown:
			return
		case <-ticker:
			opts.Scheduler.processScheduledEvents()
		}
	}
}
