package events

import (
	"context"
	"runtime"

	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/utils"
	"golang.org/x/sync/semaphore"
)

var (
	MaxWorkers = runtime.GOMAXPROCS(0)
	sem        = semaphore.NewWeighted(int64(MaxWorkers))
)

func Emit(ctx context.Context, event Event) error {
	logger, ok := logging.FromContext(ctx)
	if !ok {
		logger = logging.NewLogger("event bus: ")
	}

	sched, ok := FromContext(ctx)
	if !ok {
		return utils.ErrSchedulerContext
	}

	// acquire blocks until a worker is free - if we get an error here,
	// something went wrong
	if err := sem.Acquire(ctx, 1); err != nil {
		logger.Printf("failed to acquire semaphore: %v", err)
		return err
	}

	go func(e Event) {
		defer sem.Release(1)
		sched.Schedule(e)
	}(event)

	return nil
}
