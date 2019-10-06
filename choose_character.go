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
		characterPink: &Object{
			X:           180,
			Y:           180,
			Key:         "character_choice_" + characterPink,
			GroundBound: true,
			Lifetime:    rand.Intn(1000),
		},
		characterBlue: &Object{
			X:           220,
			Y:           180,
			Key:         "character_choice_" + characterBlue,
			GroundBound: true,
			Lifetime:    rand.Intn(1000),
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
	objects = append(
		objects,
		Object{
			X:   90,
			Y:   120,
			Key: "choose_character_screen_heading",
		},
		Object{
			X:   145,
			Y:   160,
			Key: "choose_character_screen_1",
		},
		Object{
			X:   230,
			Y:   160,
			Key: "choose_character_screen_2",
		},
	)
	return objects
}

func (cc *chooseCharacter) InvokeKeyEvent(event KeyEvent) {
	if event.Type != KeyPress {
		return
	}
	if _, ok := characters[event.Key]; ok {
		cc.choice.Set(event.Key)
		cc.transition("playing")
		return
	}
}

// characterChoice stores the character the player chose.
type characterChoice struct {
	character string
}

func (choice *characterChoice) Get() string {
	return choice.character
}

func (choice *characterChoice) Set(character string) {
	choice.character = character
}

const (
	characterPink = "1"
	characterBlue = "2"
)

var characters = map[string]struct{}{
	characterPink: {},
	characterBlue: {},
}
