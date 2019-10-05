package ld45

import (
	"math"
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
		},
	}
}

func (playing *playing) Tick(ms int) {
	playing.player.lifetime += ms
	playing.player.rotation += turnSpeed * float64(playing.player.turning()) * float64(ms) / 1000
	playing.player.x += float64(playing.player.moving()) * moveSpeed * math.Sin(playing.player.rotation) * float64(ms) / 1000
	playing.player.y += float64(playing.player.moving()) * moveSpeed * -math.Cos(playing.player.rotation) * float64(ms) / 1000
}

func (playing *playing) Objects() []Object {
	objects := make([]Object, 0)
	objects = append(objects, playing.player.ToObjects()...)
	for i := range playing.interactibles {
		objects = append(objects, playing.interactibles[i].ToObjects(playing.player)...)
	}
	return objects
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
	Tick(ms int)
	ToObjects(camera) []Object
}

type bush struct {
	x float64
	y float64
}

func (b *bush) Tick(ms int) {}

func (b *bush) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, b.x, b.y)
	return []Object{
		{
			X:        x,
			Y:        y,
			Key:      "bush_2_berries",
			Lifetime: 0,
		},
	}
}

func calculateScreenPosition(cam camera, ox, oy float64) (x int, y int) {
	cx, cy := cam.Position()
	dx, dy := cx-ox, cy-oy
	rx := -dx*math.Cos(cam.Rotation()) - dy*math.Sin(cam.Rotation())
	ry := dx*math.Sin(cam.Rotation()) - dy*math.Cos(cam.Rotation())
	return int(rx) + playerX, int(ry) + playerY
}
