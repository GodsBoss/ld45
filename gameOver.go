package ld45

import "math"

type gameOver struct {
	transitioner
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
	over.objects = append(
		over.objects,
		Object{
			X:   125,
			Y:   280,
			Key: "game_over_back_to_title_hint",
		},
	)
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
	over.addAroundItems(
		weightedRandomStrings{
			"bush_3_berries":  5,
			"tree_3":          5,
			"heart_full":      5,
			"item_diamond":    5,
			"furnace_burning": 5,
			"item_gold_ingot": 5,
			"rock_gold_ore":   5,
			"rock_diamond":    5,
		},
	)
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
	over.addAroundItems(
		weightedRandomStrings{
			"bush_0_berries": 5,
			"tree_1":         5,
			"item_rock":      5,
			"heart_empty":    5,
			"furnace_off":    5,
			"rock_stone":     5,
		},
	)
}

func (over *gameOver) addAroundItems(wsr weightedRandomStrings) {
	objs := make([]Object, gameOverItemsAroundCount)
	for i := 0; i < gameOverItemsAroundCount; i++ {
		angle := 2.0 * math.Pi * (float64(i) / float64(gameOverItemsAroundCount))
		x, y := 200.0+gameOveritemsRadius*math.Sin(angle), 150.0+gameOveritemsRadius*math.Cos(angle)
		objs[i] = Object{
			X:           x,
			Y:           y,
			Key:         wsr.Random(),
			GroundBound: true,
		}
	}
	over.objects = append(over.objects, objs...)
}

const gameOverItemsAroundCount = 24
const gameOveritemsRadius = 120.0

func (over *gameOver) Tick(ms int) {}

func (over *gameOver) Objects() []Object {
	return over.objects
}

func (over *gameOver) InvokeKeyEvent(event KeyEvent) {
	if event.Type == KeyPress && event.Key == "b" {
		over.transition("title")
	}
}
