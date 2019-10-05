package ld45

type rock struct {
	x float64
	y float64

	key string
}

func (r *rock) Position() (float64, float64) {
	return r.x, r.y
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
