package ld45

type ChooseCharacter struct{}

func (cc *ChooseCharacter) ID() string {
	return "choose_character"
}

func (cc *ChooseCharacter) Tick(ms int) {}

func (cc *ChooseCharacter) Objects() []Object {
	return make([]Object, 0)
}

func (cc *ChooseCharacter) InvokeKeyEvent(event KeyEvent) {}
