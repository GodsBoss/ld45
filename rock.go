package ld45

type rock struct {
	x float64
	y float64

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
