package ld45

type title struct {
	transitioner
}

func (title *title) ID() string {
	return "title"
}

func (title *title) Init() {}

func (title *title) Tick(ms int) {}

func (title *title) Objects() []Object {
	return []Object{
		{
			X:   30,
			Y:   50,
			Key: "title_screen_title",
		},
	}
}

func (title *title) InvokeKeyEvent(event KeyEvent) {
	if event.Type == KeyPress && event.Key == "b" {
		title.transition("choose_character")
	}
}
