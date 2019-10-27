package ld45

type Game struct {
	currentState State

	states          map[string]State
	title           *title
	playing         *playing
	gameOver        *gameOver
	chooseCharacter *chooseCharacter
	eventLock       lock
}

// lock abstracts away sync.Mutex. This is useful because JS in the browser
// does not need locking - it is single-threaded (except for webworkers).
type lock interface {
	Lock()
	Unlock()
}

type GameConfiguration struct{}

func NewGame(config *GameConfiguration) *Game {
	choice := &characterChoice{}
	game := &Game{}
	result := &playResult{}
	game.states = map[string]State{
		"title": &title{
			transitioner: game,
		},
		"playing": &playing{
			transitioner: game,
			choice:       choice,
			result:       result,
		},
		"game_over": &gameOver{
			transitioner: game,
			choice:       choice,
			result:       result,
		},
		"choose_character": &chooseCharacter{
			transitioner: game,
			choice:       choice,
		},
	}
	game.transition("title")
	game.eventLock = createLock()
	return game
}

func (game *Game) StateID() string {
	return game.currentState.ID()
}

func (game *Game) Tick(ms int) {
	game.eventLock.Lock()
	defer game.eventLock.Unlock()
	game.currentState.Tick(ms)
}

func (game *Game) Objects() []Object {
	game.eventLock.Lock()
	defer game.eventLock.Unlock()
	return game.currentState.Objects()
}

func (game *Game) InvokeKeyEvent(event KeyEvent) {
	game.eventLock.Lock()
	defer game.eventLock.Unlock()
	game.currentState.InvokeKeyEvent(event)
}

func (game *Game) transition(nextStateKey string) {
	if nextState, ok := game.states[nextStateKey]; ok {
		nextState.Init()
		game.currentState = nextState
	}
}

type transitioner interface {
	transition(nextStateKey string)
}

type State interface {
	ID() string
	Init()
	Tick(ms int)
	Objects() []Object
	InvokeKeyEvent(event KeyEvent)
}
