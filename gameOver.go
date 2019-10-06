package ld45

type gameOver struct {
	choice *characterChoice
	result *playResult

	objects []Object
}

func (over *gameOver) ID() string {
	return "game_over"
}

func (over *gameOver) Init() {
	if over.result.IsVictory() {
		over.initVictoryObjects()
	} else {
		over.initDefeatObjects()
	}
}

func (over *gameOver) initVictoryObjects() {
	over.objects = []Object{
		{
			X:   190,
			Y:   130,
			Key: "game_over_crown",
		},
		{
			X:   157,
			Y:   150,
			Key: "game_over_victory_header",
		},
	}
}

func (over *gameOver) initDefeatObjects() {
	over.objects = []Object{
		{
			X:   192,
			Y:   130,
			Key: "game_over_gravestone",
		},
		{
			X:   140,
			Y:   155,
			Key: "game_over_defeat_header",
		},
	}
}

func (over *gameOver) Tick(ms int) {}

func (over *gameOver) Objects() []Object {
	return over.objects
}

func (over *gameOver) InvokeKeyEvent(event KeyEvent) {}
