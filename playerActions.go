package ld45

// playerActions are those that can be invoked if there is no entity to interact with.
type playerActions struct {
	nopOnPlayerContact
	positionPartial
	nopTick
}

func (pa *playerActions) ID() interactibleID {
	return "player_actions"
}

func (pa *playerActions) Interactions() []interaction {
	return []interaction{
		newSimpleInteraction(
			"interaction_eating_berry",
			false,
			func(p *player) bool {
				return !p.saturation.IsMaximum() && p.inventory.has(itemBerry, 1)
			},
			func(_ int, p *playing) {
				p.player.saturation.Inc(saturationPerBerry)
				p.player.inventory.add(itemBerry, -1)
			},
		),
		newSimpleInteraction(
			"interaction_plant_tree",
			false,
			func(p *player) bool {
				return p.inventory.has(itemSapling, 1)
			},
			func(_ int, p *playing) {
				p.player.inventory.add(itemSapling, -1)
				plX, plY := p.player.Position()
				tx, ty := relativePosition(plX, plY, 0, -15.0, p.player.Rotation())
				p.interactibles.add(
					newTree(tx, ty, 1),
				)
			},
		),
		newSimpleInteraction(
			"interaction_put_workbench",
			false,
			func(p *player) bool {
				return p.inventory.has(itemWood, 4)
			},
			func(_ int, p *playing) {
				p.player.inventory.add(itemWood, -4)
				plX, plY := p.player.Position()
				wx, wy := relativePosition(plX, plY, 0, -20.0, p.player.Rotation())
				p.interactibles.add(
					newWorkbench(wx, wy),
				)
			},
		),
		newSimpleInteraction(
			"interaction_put_furnace",
			false,
			func(p *player) bool {
				return p.inventory.has(itemFurnace, 1)
			},
			func(_ int, p *playing) {
				p.player.inventory.add(itemFurnace, -1)
				plX, plY := p.player.Position()
				fx, fy := relativePosition(plX, plY, 0, -20.0, p.player.Rotation())
				p.interactibles.add(
					newFurnace(p, fx, fy),
				)
			},
		),
	}
}

const saturationPerBerry = 2

func (pa *playerActions) ToObjects(camera) []Object {
	return make([]Object, 0)
}
