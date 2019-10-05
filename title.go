package ld45

type Title struct{}

func (title *Title) ID() string {
	return "title"
}

func (title *Title) Tick() {}

func (title *Title) Objects() []Object {
	return make([]Object, 0)
}
