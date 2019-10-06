package ld45

import "math/rand"

type interactionHub struct {
	playing *playing
}

func (hub *interactionHub) Tick(ms int) {}

func (hub *interactionHub) Objects() []Object {
	return make([]Object, 0)
}

// getInteractingInteractible returns the interactible which is in the interaction
// area of the player and is the nearest possible interactible.
func (hub *interactionHub) getInteractingInteractible() (int, interactible) {
	var currentID int
	var currentInteractible interactible
	var currentDistance float64
	hub.playing.interactibles.each(
		func(id int, i interactible) {
			ix, iy := i.Position()
			x, y := calculateScreenPosition(hub.playing.player, ix, iy)
			if !inInteractionArea(x, y) {
				return
			}
			distance := distanceToPlayer(x, y)
			if currentInteractible != nil && distance > currentDistance {
				return
			}
			currentID, currentInteractible, currentDistance = id, i, distance
		},
	)
	return currentID, currentInteractible
}

func (hub *interactionHub) playerInteractsDirectly() {
	index, i := hub.getInteractingInteractible()
	if i == nil {
		return
	}
	candidates := filterInteractions(i.Interactions(), isDirect)
	if len(candidates) == 0 {
		return
	}
	candidates[rand.Intn(len(candidates))].invoke(index, hub.playing)
}

func (hub *interactionHub) playerInteractsIndirectly() {
	index, i := hub.getInteractingInteractible()
	if i == nil {
		return
	}
	candidates := filterInteractions(i.Interactions(), isIndirect)
	if len(candidates) == 0 {
		return
	}
	if playerChoiceInteractionID, ok := hub.playing.player.chosenInteraction[i.ID()]; ok {
		for _, candidate := range candidates {
			if playerChoiceInteractionID == candidate.ID() {
				candidate.invoke(index, hub.playing)
				return
			}
		}
	}
	hub.playing.player.chosenInteraction[i.ID()] = candidates[0].ID()
	candidates[0].invoke(index, hub.playing)
}

func (hub *interactionHub) changeIndirectPlayerChoice(direction int) {
	_, i := hub.getInteractingInteractible()
	if i == nil {
		return
	}
	interactions := filterInteractions(i.Interactions(), isIndirect)
	if len(interactions) == 0 {
		return
	}
	ids := extractInteractionIDs(interactions)
	currentID, ok := hub.playing.player.chosenInteraction[i.ID()]

	// If player had not made a choice before, just use the first ID.
	if !ok {
		hub.playing.player.chosenInteraction[i.ID()] = ids[0]
		return
	}

	// Player already made a choice earlier. Find it and switch to another ID.
	for idIndex := range ids {
		if ids[idIndex] == currentID {
			nextIndex := idIndex + direction
			if nextIndex < 0 {
				nextIndex = len(ids) - 1
			}
			if nextIndex > len(ids)-1 {
				nextIndex = 0
			}
			hub.playing.player.chosenInteraction[i.ID()] = ids[nextIndex]
			return
		}
	}

	// Player already made a choice earlier, but that choice is no longer accessible.
	// Just use the first possible choice.
	hub.playing.player.chosenInteraction[i.ID()] = ids[0]
}
