package ld45

import (
	"fmt"
)

type bush struct {
	positionPartial
	nopOnPlayerContact

	growth       intProperty
	fluentGrowth float64
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
			Key:         fmt.Sprintf("bush_%d_berries", b.growth.current),
			Lifetime:    0,
			GroundBound: true,
		},
	}
}

func (b *bush) Interactions() []interaction {
	return make([]interaction, 0)
}
