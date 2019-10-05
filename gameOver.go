package ld45

type gameOver struct {
	choice *characterChoice
}

func (over *gameOver) ID() string {
	return "game_over"
}

func (over *gameOver) Tick(ms int) {}

func (over *gameOver) Objects() []Object {
	return make([]Object, 0)
}

func (over *gameOver) InvokeKeyEvent(event KeyEvent) {}
