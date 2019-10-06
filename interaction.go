package ld45

type interaction interface {
	ID() string
	possible(*player) bool
	invoke(id int, p *playing)
}

type simpleInteraction struct {
	id           string
	possibleFunc func(*player) bool
	invokeFunc   func(int, *playing)
}

func newSimpleInteraction(id string, possibleFunc func(*player) bool, invokeFunc func(id int, p *playing)) *simpleInteraction {
	return &simpleInteraction{
		id:           id,
		possibleFunc: possibleFunc,
		invokeFunc:   invokeFunc,
	}
}

func (si *simpleInteraction) ID() string {
	return si.id
}

func (si *simpleInteraction) possible(p *player) bool {
	return si.possibleFunc(p)
}

func (si *simpleInteraction) invoke(id int, p *playing) {
	si.invokeFunc(id, p)
}
