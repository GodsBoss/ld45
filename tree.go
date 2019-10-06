package ld45

import (
	"fmt"
	"math/rand"
)

type tree struct {
	positionPartial
	nopOnPlayerContact

	growth       intProperty
	fluentGrowth float64
	health       float64
}

func newTree(x, y float64, initialGrowth int) *tree {
	t := &tree{
		positionPartial: createPositionPartial(x, y),
		growth: intProperty{
			minimum: 1,
			maximum: 3,
		},
	}
	t.growth.Set(initialGrowth)
	t.healthByGrowth()
	return t
}

func (t *tree) ID() string {
	return "tree"
}

func (t *tree) Tick(ms int) {
	if !t.growth.IsMaximum() {
		t.fluentGrowth += float64(ms) / 1000.0
		if t.fluentGrowth >= treeGrowCost {
			t.growth.Inc(1)
			t.fluentGrowth = 0
			t.healthByGrowth()
		}
	}
}

func (t *tree) healthByGrowth() {
	t.health = healthPerSize * float64(t.growth.current)
}

const treeGrowCost = 60.0
const healthPerSize = 5.0

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

func (t *tree) Interactions() []interaction {
	return []interaction{
		newSimpleInteraction(
			"chop_tree",
			possibleAlways,
			func(id int, p *playing) {
				t.health -= float64(p.player.equipment[toolAxe]+1) * 2.0
				if t.health <= 0 {
					sx, sy := randomPositionAround(t.x, t.y, 10.0, 20.0)
					p.interactibles.add(
						itemSapling.New(sx, sy),
					)
					if t.growth.IsMaximum() {
						count := rand.Intn(3) + 1
						for i := 0; i < count; i++ {
							wx, wy := randomPositionAround(t.x, t.y, 10.0, 25.0)
							p.interactibles.add(
								itemWood.New(wx, wy),
							)
						}
					}
					p.interactibles.remove(id)
				}
			},
		),
	}
}
