package ld45

type Playing struct{}

func (playing *Playing) ID() string {
	return "playing"
}

func (playing *Playing) Tick() {}

func (playing *Playing) Objects() []Object {
	return make([]Object, 0)
}
