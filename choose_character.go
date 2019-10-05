package ld45

import (
	"math/rand"
)

type chooseCharacter struct {
	transitioner
	choice     *characterChoice
	characters map[string]*Object
}

func (cc *chooseCharacter) ID() string {
	return "choose_character"
}

func (cc *chooseCharacter) Init() {
	cc.characters = map[string]*Object{
		"1": &Object{
			X:        180,
			Y:        50,
			Key:      "character_choice_1",
			Lifetime: rand.Intn(1000),
		},
		"2": &Object{
			X:        220,
			Y:        50,
			Key:      "character_choice_2",
			Lifetime: rand.Intn(1000),
		},
	}
}

func (cc *chooseCharacter) Tick(ms int) {
	for key := range cc.characters {
		cc.characters[key].Lifetime += ms/(rand.Intn(3)+1) + rand.Intn(10)
	}
}

func (cc *chooseCharacter) Objects() []Object {
	objects := make([]Object, len(cc.characters))
	index := 0
	for key := range cc.characters {
		objects[index] = *(cc.characters[key])
		index++
	}
	return objects
}

func (cc *chooseCharacter) InvokeKeyEvent(event KeyEvent) {
	if event.Type != KeyPress {
		return
	}
	if event.Key == "1" {
		cc.choice.Set(characterPink)
		cc.transition("playing")
		return
	}
	if event.Key == "2" {
		cc.choice.Set(characterBlue)
		cc.transition("playing")
		return
	}
}

// characterChoice stores the character the player chose.
type characterChoice struct {
	character int
}

func (choice *characterChoice) Set(character int) {
	choice.character = character
}

const (
	characterPink = iota
	characterBlue
)
