package events

import (
	"context"
	"time"
)

type EventType int

type Event interface {
	Process(ctx context.Context)
	Duration(ctx context.Context) (time.Duration, error)
	String() string
}
