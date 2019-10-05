package ld45

type interaction interface {
	possible(*player) bool
	invoke(id int, p *playing)
}

type simpleInteraction struct {
	possibleFunc func(*player) bool
	invokeFunc   func(int, *playing)
}

func newSimpleInteraction(possibleFunc func(*player) bool, invokeFunc func(id int, p *playing)) *simpleInteraction {
	return &simpleInteraction{
		possibleFunc: possibleFunc,
		invokeFunc:   invokeFunc,
	}
}

func (si *simpleInteraction) possible(p *player) bool {
	return si.possibleFunc(p)
}

func (si *simpleInteraction) invoke(id int, p *playing) {
	si.invokeFunc(id, p)
}
