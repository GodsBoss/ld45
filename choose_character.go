package ld45

type ChooseCharacter struct {
	choice *characterChoice
}

func (cc *ChooseCharacter) ID() string {
	return "choose_character"
}

func (cc *ChooseCharacter) Tick(ms int) {}

func (cc *ChooseCharacter) Objects() []Object {
	return make([]Object, 0)
}

func (cc *ChooseCharacter) InvokeKeyEvent(event KeyEvent) {}

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
