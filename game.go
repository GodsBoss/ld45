package ld45

type Game struct {
	currentState State

	states          map[string]State
	title           *Title
	playing         *Playing
	gameOver        *GameOver
	chooseCharacter *ChooseCharacter
}

func NewGame() *Game {
	choice := &characterChoice{}
	game := &Game{
		states: map[string]State{
			"title": &Title{},
			"playing": &Playing{
				choice: choice,
			},
			"game_over": &GameOver{
				choice: choice,
			},
			"choose_character": &ChooseCharacter{
				choice: choice,
			},
		},
	}
	game.transition("title")
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

func (game *Game) transition(nextStateKey string) {
	if nextState, ok := game.states[nextStateKey]; ok {
		game.currentState = nextState
	}
}

type transitioner interface {
	transition(nextStateKey string)
}

type State interface {
	ID() string
	Tick(ms int)
	Objects() []Object
	InvokeKeyEvent(event KeyEvent)
}
