package ld45

type interactible interface {
	Interactions() []interaction
	Position() (float64, float64)
	Tick(ms int)
	ToObjects(camera) []Object
}

type positionPartial struct {
	x float64
	y float64
}

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
