package ld45

type title struct {
	transitioner
}

func (title *title) ID() string {
	return "title"
}

func (title *title) Tick(ms int) {}

func (title *title) Objects() []Object {
	return make([]Object, 0)
}

func (title *title) InvokeKeyEvent(event KeyEvent) {
	if event.Type == KeyPress && event.Key == "b" {
		title.transition("choose_character")
	}
}
