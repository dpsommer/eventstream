package characters

import (
	"context"

	"github.com/dpsommer/eventstream/internal/events"
	"github.com/dpsommer/eventstream/internal/logging"
	"github.com/dpsommer/eventstream/internal/regions"
)

// the observer should:
// * watch over actor objects
// * track state/determine behaviour
// * schedule events
// * lock/unlock instances (specific behaviours?)
// * keep track of relationships between actors
// * prune 'dead' actors (?)
// *

type Observer struct {
	characters []*Character
	ctx        context.Context
}

func NewObserver(ctx context.Context) *Observer {
	return &Observer{
		characters: []*Character{},
		ctx:        ctx,
	}
}

func (o *Observer) AddCharacter(loc regions.Location) {
	s := chooseSex()
	g := chooseGender(s)
	o.characters = append(o.characters, &Character{
		name:     generateName(g),
		surname:  generateSurname(),
		location: loc,
		gender:   g,
		sex:      s,
	})
}

func (o *Observer) ProcessCharaterBehaviour() {
	logger, ok := logging.FromContext(o.ctx)
	if !ok {
		logger = logging.NewLogger("observer: ")
	}

	for _, c := range o.characters {
		// TODO: character state machine?
		// need a way to determine character behaviour
		if !c.IsMoving(o.ctx) {
			var dest regions.Location

			switch c.location {
			case regions.Town:
				dest = regions.Forest
			default:
				dest = regions.Town
			}

			err := events.Emit(o.ctx, &events.MovementEvent{
				Destination: dest,
				Movable:     c,
			})
			if err != nil {
				logger.Printf("%v\n", err.Error())
			}

			c.isMoving = true
		}
	}
}
