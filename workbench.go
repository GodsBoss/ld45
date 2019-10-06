package ld45

type workbench struct {
	nopOnPlayerContact
	positionPartial
	nopTick
}

func (wb *workbench) ID() string {
	return "workbench"
}

func (wb *workbench) Interactions() []interaction {
	return make([]interaction, 0)
}

func (wb *workbench) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, wb.x, wb.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         "workbench",
			GroundBound: true,
		},
	}
}
