package characters

import (
	"context"
	"sync"

	"github.com/dpsommer/eventstream/internal/regions"
)

type Character struct {
	Name     string
	Location regions.Location

	sync.Mutex
}

func NewCharacter(ctx context.Context, name string) *Character {
	return &Character{
		Name:     name,
		Location: regions.Town,
	}
}
