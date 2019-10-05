package ld45

type playing struct {
	choice *characterChoice
}

func (playing *playing) ID() string {
	return "playing"
}

func (playing *playing) Init() {}

func (playing *playing) Tick(ms int) {}

func (playing *playing) Objects() []Object {
	return make([]Object, 0)
}

func (playing *playing) InvokeKeyEvent(event KeyEvent) {}
