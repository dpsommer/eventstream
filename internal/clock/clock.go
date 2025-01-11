package clock

import (
	"context"
	"os"
	"time"

	"github.com/dpsommer/eventstream/internal/characters"
	"github.com/dpsommer/eventstream/internal/events"
	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/utils"
)

type WorldClockOptions struct {
	Tick     time.Duration
	Shutdown <-chan os.Signal
}

func StartWorldClock(ctx context.Context, opts *WorldClockOptions) {
	logger, ok := logging.FromContext(ctx)
	if !ok {
		logger = logging.NewLogger("clock: ")
	}

	sched, ok := events.FromContext(ctx)
	if !ok {
		logger.Printf("failed to start world clock: %v", utils.ErrSchedulerContext)
		return
	}

	observer, ok := characters.FromContext(ctx)
	if !ok {
		logger.Printf("failed to start world clock: %v", utils.ErrObserverContext)
		return
	}

	ticker := time.Tick(opts.Tick)

	for {
		select {
		case <-opts.Shutdown:
			return
		case <-ticker:
			observer.ProcessCharaterBehaviour()
			sched.ProcessScheduledEvents()
		}
	}
}
