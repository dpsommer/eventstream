package events

import (
	"container/heap"
	"context"
	"time"

	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/utils"
)

// scheduled event implements Prioritizable
type ScheduledEvent struct {
	priority int64

	Event
}

func (se ScheduledEvent) Priority() int64 { return se.priority }

type Scheduler struct {
	scheduledEvents *utils.PriorityQueue[ScheduledEvent]
	ctx             context.Context
}

func NewScheduler(ctx context.Context) *Scheduler {
	return &Scheduler{
		scheduledEvents: &utils.PriorityQueue[ScheduledEvent]{},
		ctx:             ctx,
	}
}

func (s *Scheduler) processScheduledEvents() {
	for {
		if s.scheduledEvents.Len() < 1 {
			continue
		}

		se := heap.Pop(s.scheduledEvents).(*ScheduledEvent)
		t := time.Unix(0, se.priority)

		// while the front of the minheap is after the current time, continue
		// to pop and process scheduled events. otherwise, put the event back
		// in the queue and break until the next tick
		if time.Now().After(t) {
			se.Process(s.ctx)
		} else {
			heap.Push(s.scheduledEvents, se)
			break
		}
	}
}

func (s *Scheduler) Schedule(e Event) {
	logger, ok := logging.FromContext(s.ctx)
	if !ok {
		logger, s.ctx = logging.WithContext(s.ctx, "scheduler: ")
	}

	duration, err := e.Duration(s.ctx)
	if err != nil {
		logger.Printf("failed to schedule event %q: %s\n", e.String(), err.Error())
		return
	}

	heap.Push(s.scheduledEvents, &ScheduledEvent{
		priority: time.Now().Add(duration).UnixNano(),
		Event:    e,
	})
}
