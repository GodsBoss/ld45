package ld45

import (
	"math/rand"
	"sort"
)

type title struct {
	transitioner

	things *titleThings
}

func (title *title) ID() string {
	return "title"
}

func (title *title) Init() {
	title.things = &titleThings{
		m: make(map[int]*titleThing),
		keyDistribution: weightedRandomStrings{
			"tree_1":          20,
			"tree_2":          25,
			"tree_3":          30,
			"rock_stone":      25,
			"rock_coal":       15,
			"rock_iron_ore":   10,
			"rock_gold_ore":   5,
			"rock_diamond":    1,
			"bush_1_berries":  10,
			"bush_2_berries":  15,
			"bush_3_berries":  20,
			"workbench":       5,
			"furnace_off":     3,
			"furnace_burning": 1,
		},
	}
	for i := 0; i < titleThingCount; i++ {
		title.things.addRandomTitleThing(true)
	}
}

func (title *title) Tick(ms int) {
	title.things.Tick(ms)
}

func (title *title) Objects() []Object {
	objs := append(
		title.things.Objects(),
		Object{
			X:   30,
			Y:   50,
			Key: "title_screen_title",
		},
		Object{
			X:   110,
			Y:   270,
			Key: "title_screen_b_hint",
		},
	)
	sortedObjs := Objects(objs)
	sort.Sort(sortedObjs)
	return objs
}

func (title *title) InvokeKeyEvent(event KeyEvent) {
	if event.Type == KeyPress && event.Key == "b" {
		title.transition("choose_character")
	}
}

type titleThings struct {
	m               map[int]*titleThing
	index           int
	keyDistribution weightedRandomStrings
}

func (things *titleThings) Objects() []Object {
	objs := make([]Object, len(things.m))
	index := 0
	for _, thing := range things.m {
		objs[index] = thing.Object()
		index++
	}
	return objs
}

func (things *titleThings) Tick(ms int) {
	for index := range things.m {
		things.m[index].Tick(ms)
		if things.m[index].Y > 320.0 {
			delete(things.m, index)
			things.addRandomTitleThing(false)
		}
	}
}

func (things *titleThings) add(thing *titleThing) {
	things.m[things.index] = thing
	things.index++
}

func (things *titleThings) addRandomTitleThing(randomY bool) {
	y := -10.0
	if randomY {
		y = rand.Float64()*310.0 - 10.0
	}
	things.add(
		&titleThing{
			key:      things.keyDistribution.Random(),
			lifetime: 0,
			X:        rand.Float64()*410.0 - 5.0,
			Y:        y,
		},
	)
}

type titleThing struct {
	key      string
	lifetime int
	X        float64
	Y        float64
}

func (thing *titleThing) Object() Object {
	return Object{
		X:           int(thing.X),
		Y:           int(thing.Y),
		Key:         thing.key,
		Lifetime:    thing.lifetime,
		GroundBound: true,
	}
}

func (thing *titleThing) Tick(ms int) {
	thing.lifetime += ms
	thing.Y += titleMoveSpeed * float64(ms) / 1000.0
}

const titleMoveSpeed = 10.0

const titleThingCount = 100
