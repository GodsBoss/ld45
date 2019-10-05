package ld45

type rock struct {
	positionPartial
	nopOnPlayerContact

	key string
}

func (r *rock) Tick(ms int) {}

func (r *rock) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, r.x, r.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         r.key,
			Lifetime:    0,
			GroundBound: true,
		},
	}
}

func (r *rock) Interactions() []interaction {
	return make([]interaction, 0)
}
