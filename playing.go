package ld45

import (
	"math"
	"math/rand"
	"sort"
)

type playing struct {
	choice *characterChoice

	player        *player
	interactibles []interactible
}

func (playing *playing) ID() string {
	return "playing"
}

func (playing *playing) Init() {
	playing.player = &player{
		key: playing.choice.character,
		health: &intProperty{
			maximum: 20,
			current: 20,
		},
		saturation: &intProperty{
			maximum: 20,
			current: 20,
		},
		rotation: 0,
		x:        0,
		y:        0,
	}
	playing.interactibles = []interactible{
		&bush{
			x: 0.0,
			y: -50.0,
			growth: intProperty{
				maximum: 3,
			},
		},
		&bush{
			x: -25.0,
			y: -60.0,
			growth: intProperty{
				maximum: 3,
			},
		},
		&bush{
			x: 60.0,
			y: -25.0,
			growth: intProperty{
				maximum: 3,
			},
		},
		&tree{
			x: -80.0,
			y: -80.0,
			growth: intProperty{
				minimum: 1,
				maximum: 3,
				current: 1,
			},
		},
		&tree{
			x: -120.0,
			y: 25.0,
			growth: intProperty{
				minimum: 1,
				maximum: 3,
				current: 2,
			},
		},
		&tree{
			x: 90.0,
			y: 5.0,
			growth: intProperty{
				minimum: 1,
				maximum: 3,
				current: 3,
			},
		},
		&rock{
			x:   100.0,
			y:   25.0,
			key: "rock",
		},
		&rock{
			x:   -90.0,
			y:   -75.0,
			key: "coal_ore",
		},
		&rock{
			x:   55.0,
			y:   40.0,
			key: "iron_ore",
		},
		&rock{
			x:   60.0,
			y:   -90.0,
			key: "gold_ore",
		},
	}
}

func (playing *playing) Tick(ms int) {
	playing.player.lifetime += ms
	playing.player.rotation += turnSpeed * float64(playing.player.turning()) * float64(ms) / 1000
	playing.player.x += float64(playing.player.moving()) * moveSpeed * math.Sin(playing.player.rotation) * float64(ms) / 1000
	playing.player.y += float64(playing.player.moving()) * moveSpeed * -math.Cos(playing.player.rotation) * float64(ms) / 1000
	for i := range playing.interactibles {
		playing.interactibles[i].Tick(ms)
	}
}

func (playing *playing) Objects() []Object {
	objects := make(Objects, 0)
	objects = append(objects, playing.player.ToObjects()...)
	for i := range playing.interactibles {
		objects = append(objects, playing.interactibles[i].ToObjects(playing.player)...)
	}
	sort.Sort(objects)
	return objects
}

func (playing *playing) playerInteracts() {
	interactionCandidates := make([]interaction, 0)
	for i := range playing.interactibles {
		ix, iy := playing.interactibles[i].Position()
		x, y := calculateScreenPosition(playing.player, ix, iy)
		if inInteractionArea(x, y) {
			interactionCandidates = append(interactionCandidates, playing.player.filterInteractions(playing.interactibles[i].Interactions())...)
		}
	}
	if len(interactionCandidates) == 0 {
		return
	}
	chosenInteractible := interactionCandidates[rand.Intn(len(interactionCandidates))]
	chosenInteractible.invoke()
}

func (playing *playing) InvokeKeyEvent(event KeyEvent) {
	switch event.Key {
	case "a":
		if event.Type == KeyDown {
			playing.player.turnLeft = true
		}
		if event.Type == KeyUp {
			playing.player.turnLeft = false
		}
	case "d":
		if event.Type == KeyDown {
			playing.player.turnRight = true
		}
		if event.Type == KeyUp {
			playing.player.turnRight = false
		}
	case "w":
		if event.Type == KeyDown {
			playing.player.moveForward = true
		}
		if event.Type == KeyUp {
			playing.player.moveForward = false
		}
	case "s":
		if event.Type == KeyDown {
			playing.player.moveBackward = true
		}
		if event.Type == KeyUp {
			playing.player.moveBackward = false
		}
	case "q":
		if event.Type == KeyDown {
			playing.playerInteracts()
		}
	}
}

type camera interface {
	Position() (float64, float64)
	Rotation() float64
}

type intProperty struct {
	current int
	minimum int
	maximum int
}

func (prop *intProperty) IsMinimum() bool {
	return prop.current == prop.minimum
}

func (prop *intProperty) IsMaximum() bool {
	return prop.current == prop.maximum
}

func (prop *intProperty) Dec(amount int) {
	prop.current -= amount
	if prop.current < prop.minimum {
		prop.current = prop.minimum
	}
}

func (prop *intProperty) Inc(amount int) {
	prop.current += amount
	if prop.current > prop.maximum {
		prop.current = prop.maximum
	}
}

var boolToInt = map[bool]int{
	false: 0,
	true:  1,
}

type interactible interface {
	Interactions() []interaction
	Position() (float64, float64)
	Tick(ms int)
	ToObjects(camera) []Object
}

type interaction interface {
	possible(*player) bool
	invoke()
}

func calculateScreenPosition(cam camera, ox, oy float64) (x int, y int) {
	cx, cy := cam.Position()
	dx, dy := cx-ox, cy-oy
	rx := -dx*math.Cos(cam.Rotation()) - dy*math.Sin(cam.Rotation())
	ry := dx*math.Sin(cam.Rotation()) - dy*math.Cos(cam.Rotation())
	return int(rx) + playerX, int(ry) + playerY
}
