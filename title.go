package ld45

type title struct{}

func (title *title) ID() string {
	return "title"
}

func (title *title) Tick(ms int) {}

func (title *title) Objects() []Object {
	return make([]Object, 0)
}

func (title *title) InvokeKeyEvent(event KeyEvent) {}
