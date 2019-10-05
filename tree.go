package ld45

import "fmt"

type tree struct {
	x float64
	y float64

	growth       intProperty
	fluentGrowth float64
}

func (t *tree) Position() (float64, float64) {
	return t.x, t.y
}

func (t *tree) Tick(ms int) {
	if !t.growth.IsMaximum() {
		t.fluentGrowth += float64(ms) / 1000.0
		if t.fluentGrowth >= treeGrowCost {
			t.growth.Inc(1)
			t.fluentGrowth = 0
		}
	}
}

const treeGrowCost = 60.0

func (t *tree) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, t.x, t.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         fmt.Sprintf("tree_%d", t.growth.current),
			Lifetime:    0,
			GroundBound: true,
		},
	}
}

func (t *tree) Interact() {}
