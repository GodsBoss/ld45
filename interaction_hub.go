package ld45

import "math/rand"

type interactionHub struct {
	playing *playing

	// chosenInteraction stores the interaction concerning a given interactible ID.
	chosenInteraction map[interactibleID]interactionID

	currentChoice *interactionID
}

func (hub *interactionHub) Tick(ms int) {
	_, choice := hub.getIndirectPlayerChoice()
	if choice != nil {
		id := choice.ID()
		hub.currentChoice = &id
	} else {
		hub.currentChoice = nil
	}
}

func (hub *interactionHub) Objects() []Object {
	objects := make([]Object, 0)
	if hub.currentChoice != nil {
		objects = append(
			objects,
			Object{
				X:   200,
				Y:   250,
				Key: string(*hub.currentChoice),
			},
			Object{
				X:   199,
				Y:   249,
				Key: "interaction_marker",
			},
		)
	}
	return objects
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

func (hub *interactionHub) filterForValidInteractions(interactions []interaction) []interaction {
	return filterInteractions(interactions, playerIsAble(hub.playing.player))
}

func (hub *interactionHub) playerInteractsDirectly() {
	index, i := hub.getInteractingInteractible()
	if i == nil {
		return
	}
	candidates := filterInteractions(filterInteractions(i.Interactions(), isDirect), playerIsAble(hub.playing.player))
	if len(candidates) == 0 {
		return
	}
	candidates[rand.Intn(len(candidates))].invoke(index, hub.playing)
}

func (hub *interactionHub) getIndirectPlayerChoice() (int, interaction) {
	index, i := hub.getInteractingInteractible()
	if i == nil {
		return -1, nil
	}
	candidates := filterInteractions(filterInteractions(i.Interactions(), isIndirect), playerIsAble(hub.playing.player))
	if len(candidates) == 0 {
		return -1, nil
	}
	if playerChoiceInteractionID, ok := hub.chosenInteraction[i.ID()]; ok {
		for _, candidate := range candidates {
			if playerChoiceInteractionID == candidate.ID() {
				return index, candidate
			}
		}
	}
	hub.chosenInteraction[i.ID()] = candidates[0].ID()
	return index, candidates[0]
}

func (hub *interactionHub) playerInteractsIndirectly() {
	index, chosenInteraction := hub.getIndirectPlayerChoice()
	if chosenInteraction == nil {
		return
	}
	chosenInteraction.invoke(index, hub.playing)
}

func (hub *interactionHub) changeIndirectPlayerChoice(direction int) {
	_, i := hub.getInteractingInteractible()
	if i == nil {
		return
	}
	interactions := filterInteractions(filterInteractions(i.Interactions(), isIndirect), playerIsAble(hub.playing.player))
	if len(interactions) == 0 {
		return
	}
	ids := extractInteractionIDs(interactions)
	currentID, ok := hub.chosenInteraction[i.ID()]

	// If player had not made a choice before, just use the first ID.
	if !ok {
		hub.chosenInteraction[i.ID()] = ids[0]
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
			hub.chosenInteraction[i.ID()] = ids[nextIndex]
			return
		}
	}

	// Player already made a choice earlier, but that choice is no longer accessible.
	// Just use the first possible choice.
	hub.chosenInteraction[i.ID()] = ids[0]
}
