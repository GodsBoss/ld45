package ld45

import (
	"math/rand"
	"strconv"
)

type bush struct {
	positionPartial
	nopOnPlayerContact
	storedInteractions

	growth       intProperty
	fluentGrowth float64
}

func newBush(x, y float64, initialGrowth int) *bush {
	b := &bush{
		positionPartial: createPositionPartial(x, y),
		growth: intProperty{
			maximum: 3,
			current: initialGrowth,
		},
		fluentGrowth: rand.Float64() * berryCost * 0.25,
	}
	b.interactions = []interaction{
		newSimpleInteraction(
			"interaction_picking_berry",
			directInteraction,
			func(_ *player) bool {
				return b.growth.current > 0
			},
			func(id int, p *playing) {
				itemX, itemY := randomPositionAround(b.x, b.y, 10.0, 20.0)
				p.interactibles.add(itemBerry.New(itemX, itemY))
				b.growth.Dec(1)
			},
		),
	}
	return b
}

func (b *bush) ID() interactibleID {
	return "bush"
}

func (b *bush) Tick(ms int) {
	if !b.growth.IsMaximum() {
		b.fluentGrowth += float64(ms) / 1000.0
		if b.fluentGrowth > berryCost {
			b.growth.Inc(1)
			b.fluentGrowth = 0.0
		}
	}
}

const berryCost = 30.0

func (b *bush) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, b.x, b.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         "bush_" + strconv.Itoa(b.growth.current) + "_berries",
			Lifetime:    0,
			GroundBound: true,
		},
	}
}
