package ld45

type Game struct {
	currentState State

	title    *Title
	playing  *Playing
	gameOver *GameOver
}

func NewGame() *Game {
	game := &Game{
		title:    &Title{},
		playing:  &Playing{},
		gameOver: &GameOver{},
	}
	game.currentState = game.title
	return game
}

func (game *Game) StateID() string {
	return game.currentState.ID()
}

func (game *Game) Tick() {
	game.currentState.Tick()
}

func (game *Game) Objects() []Object {
	return game.currentState.Objects()
}

type State interface {
	ID() string
	Tick()
	Objects() []Object
}
