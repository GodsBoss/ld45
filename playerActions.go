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
	}
}

const saturationPerBerry = 2

func (pa *playerActions) ToObjects(camera) []Object {
	return make([]Object, 0)
}
