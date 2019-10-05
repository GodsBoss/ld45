package ld45

type chooseCharacter struct {
	choice *characterChoice
}

func (cc *chooseCharacter) ID() string {
	return "choose_character"
}

func (cc *chooseCharacter) Tick(ms int) {}

func (cc *chooseCharacter) Objects() []Object {
	return make([]Object, 0)
}

func (cc *chooseCharacter) InvokeKeyEvent(event KeyEvent) {}

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
