package ld45

type bush struct {
	x float64
	y float64
}

func (b *bush) Tick(ms int) {}

func (b *bush) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, b.x, b.y)
	return []Object{
		{
			X:        x,
			Y:        y,
			Key:      "bush_2_berries",
			Lifetime: 0,
		},
	}
}
