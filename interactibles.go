package ld45

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
