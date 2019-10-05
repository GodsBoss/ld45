package ld45

type furnace struct {
	nopOnPlayerContact
	positionPartial

	lifetime int
}

func (furn *furnace) Interactions() []interaction {
	return make([]interaction, 0)
}

func (furn *furnace) Tick(ms int) {
	furn.lifetime += ms
}

func (furn *furnace) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, furn.x, furn.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         "furnace_off",
			GroundBound: true,
			Lifetime:    furn.lifetime,
		},
	}
}
