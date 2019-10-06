package ld45

type interactible interface {
	ID() interactibleID
	Interactions() []interaction
	OnPlayerContact(id int, p *playing)
	Position() (float64, float64)
	Tick(ms int)
	ToObjects(camera) []Object
}

type interactibleID string

type noInteractions struct{}

func (n noInteractions) Interactions() []interaction {
	return make([]interaction, 0)
}

type nopOnPlayerContact struct{}

func (contact nopOnPlayerContact) OnPlayerContact(_ int, _ *playing) {}

type positionPartial struct {
	x float64
	y float64
}

type nopTick struct{}

func (tick nopTick) Tick(ms int) {}

func createPositionPartial(x, y float64) positionPartial {
	return positionPartial{
		x: x,
		y: y,
	}
}

func (p positionPartial) Position() (float64, float64) {
	return p.x, p.y
}

type interactibles struct {
	lastID int
	m      map[int]interactible
}

func newInteractibles() *interactibles {
	return &interactibles{
		m: make(map[int]interactible),
	}
}

func (is *interactibles) add(i interactible) int {
	is.lastID++
	is.m[is.lastID] = i
	return is.lastID
}

func (is *interactibles) remove(id int) {
	delete(is.m, id)
}

func (is *interactibles) each(f func(id int, i interactible)) {
	for id := range is.m {
		f(id, is.m[id])
	}
}
