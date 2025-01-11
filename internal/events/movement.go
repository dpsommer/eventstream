package events

import (
	"context"
	"fmt"
	"time"

	"github.com/dpsommer/eventstream/internal/regions"
	"github.com/dpsommer/eventstream/internal/utils"
)

type Movable interface {
	Location(ctx context.Context) regions.Location
	IsMoving(ctx context.Context) bool
	Move(ctx context.Context, to regions.Location) error
}

type MovementEvent struct {
	Destination regions.Location
	Movable
}

func (e *MovementEvent) Process(ctx context.Context) {
	e.Move(ctx, e.Destination)
}

func (e *MovementEvent) Duration(ctx context.Context) (time.Duration, error) {
	m, ok := regions.FromContext(ctx)
	if !ok {
		return 0, utils.ErrMapContext
	}

	_, d, err := m.Distance(e.Location(ctx), e.Destination)
	if err != nil {
		return 0, fmt.Errorf("failed to determine movement duration: %w", err)
	}

	return time.Duration(d) * time.Second, nil
}

func (e *MovementEvent) String() string {
	return fmt.Sprintf("move to %s", e.Destination)
}
