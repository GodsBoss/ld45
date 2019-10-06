package ld45

import "sort"

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

func filterInteractions(interactions []interaction, predicate func(interaction) bool) []interaction {
	result := make([]interaction, 0)
	for i := range interactions {
		if predicate(interactions[i]) {
			result = append(result, interactions[i])
		}
	}
	return result
}

func isDirect(i interaction) bool {
	return i.IsDirect()
}

func isIndirect(i interaction) bool {
	return !i.IsDirect()
}

func playerIsAble(p *player) func(interaction) bool {
	return func(i interaction) bool {
		return i.possible(p)
	}
}

func extractInteractionIDs(interactions []interaction) interactionIDs {
	result := make(interactionIDs, len(interactions))
	for i := range interactions {
		result[i] = interactions[i].ID()
	}
	sort.Sort(result)
	return result
}

type interactionIDs []interactionID

func (ids interactionIDs) Len() int {
	return len(ids)
}

func (ids interactionIDs) Swap(i, j int) {
	ids[i], ids[j] = ids[j], ids[i]
}

func (ids interactionIDs) Less(i, j int) bool {
	return ids[i] < ids[j]
}
