package characters

import (
	"context"
	"fmt"

	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/regions"
)

type Character struct {
	name     string
	surname  string
	location regions.Location
	isMoving bool

	gender
	sex
}

func NewCharacter(ctx context.Context, name string) *Character {
	return &Character{
		name:     name,
		location: regions.Town,
	}
}

// implement the Movable interface
func (c *Character) Location(ctx context.Context) regions.Location {
	return c.location
}

func (c *Character) IsMoving(ctx context.Context) bool {
	return c.isMoving
}

func (c *Character) Move(ctx context.Context, dest regions.Location) error {
	logger, ok := logging.FromContext(ctx)
	if !ok {
		logger = logging.NewLogger(fmt.Sprintf("%s: ", c.name))
	}

	logger.Printf("%s moved from %s to %s\n", c.name, c.location, dest)

	c.location = dest
	c.isMoving = false

	return nil
}
