package ld45

type interaction interface {
	ID() interactionID
	IsDirect() bool
	possible(*player) bool
	invoke(id int, p *playing)
}

const (
	directInteraction   = true
	indirectInteraction = false
)

type interactionID string

type simpleInteraction struct {
	id           interactionID
	direct       bool
	possibleFunc func(*player) bool
	invokeFunc   func(int, *playing)
}

func newSimpleInteraction(id interactionID, direct bool, possibleFunc func(*player) bool, invokeFunc func(id int, p *playing)) *simpleInteraction {
	return &simpleInteraction{
		id:           id,
		direct:       direct,
		possibleFunc: possibleFunc,
		invokeFunc:   invokeFunc,
	}
}

func (si *simpleInteraction) ID() interactionID {
	return si.id
}

func (si *simpleInteraction) IsDirect() bool {
	return si.direct
}

func (si *simpleInteraction) possible(p *player) bool {
	return si.possibleFunc(p)
}

func (si *simpleInteraction) invoke(id int, p *playing) {
	si.invokeFunc(id, p)
}
