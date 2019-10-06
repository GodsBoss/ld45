package ld45

type rock struct {
	positionPartial
	nopOnPlayerContact

	key rockType
}

type rockType string

const (
	rockStone   rockType = "rock_stone"
	rockCoal    rockType = "rock_coal"
	rockIronOre rockType = "rock_iron_ore"
	rockGoldOre rockType = "rock_gold_ore"
	rockDiamond rockType = "rock_diamond"
)

func (r *rock) ID() interactibleID {
	return interactibleID(r.key)
}

func (r *rock) Tick(ms int) {}

func (r *rock) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, r.x, r.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         string(r.key),
			Lifetime:    0,
			GroundBound: true,
		},
	}
}

func (r *rock) Interactions() []interaction {
	return make([]interaction, 0)
}
