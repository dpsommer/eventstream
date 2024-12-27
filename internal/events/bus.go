package events

import (
	"context"
	"runtime"

	"github.com/dpsommer/eventstream/internal/logging"
	"golang.org/x/sync/semaphore"
)

var (
	MaxWorkers = runtime.GOMAXPROCS(0)
	sem        = semaphore.NewWeighted(int64(MaxWorkers))
)

func Emit(ctx context.Context, event Event) error {
	logger, ok := logging.FromContext(ctx)
	if !ok {
		logger, ctx = logging.WithContext(ctx, "event worker: ")
	}

	sched, ok := FromContext(ctx)
	if !ok {
		sched = NewScheduler(ctx)
	}

	// acquire blocks until a worker is free - if we get an error here,
	// something went wrong
	if err := sem.Acquire(ctx, 1); err != nil {
		logger.Printf("failed to acquire semaphore: %v", err)
		return err
	}

	go func(e Event) {
		defer sem.Release(1)
		// FIXME: locking here means that the event pool can get completely
		// blocked if objects have multiple blocking events fire. where is the
		// correct place to handle this synchronous behaviour? is there a
		// better approach than just slapping mutexes on each actor?
		e.PreProcess(ctx)
		sched.Schedule(e)
	}(event)

	return nil
}
