package ld45

type Game struct {
	currentState State

	title           *Title
	playing         *Playing
	gameOver        *GameOver
	chooseCharacter *ChooseCharacter
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

func (game *Game) Tick(ms int) {
	game.currentState.Tick(ms)
}

func (game *Game) Objects() []Object {
	return game.currentState.Objects()
}

func (game *Game) InvokeKeyEvent(event KeyEvent) {
	game.currentState.InvokeKeyEvent(event)
}

type State interface {
	ID() string
	Tick(ms int)
	Objects() []Object
	InvokeKeyEvent(event KeyEvent)
}
