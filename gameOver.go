package ld45

type GameOver struct{}

func (over *GameOver) ID() string {
	return "game_over"
}

func (over *GameOver) Tick(ms int) {}

func (over *GameOver) Objects() []Object {
	return make([]Object, 0)
}
