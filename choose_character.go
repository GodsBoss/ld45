package ld45

type chooseCharacter struct {
	transitioner
	choice *characterChoice
}

func (cc *chooseCharacter) ID() string {
	return "choose_character"
}

func (cc *chooseCharacter) Tick(ms int) {}

func (cc *chooseCharacter) Objects() []Object {
	return make([]Object, 0)
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
