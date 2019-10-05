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

type player struct {
	key string

	lifetime int

	health     *intProperty
	saturation *intProperty

	// rotation is the player's rotation. Zero means "up".
	rotation float64

	// turnLeft and turnRight determine wether the player attempts to move left and/or right.
	turnLeft  bool
	turnRight bool

	// moveForward and moveBackward determine wether the player attempts to move forward and/or backward.
	moveForward  bool
	moveBackward bool

	// x and y are the player's coordinates.
	x float64
	y float64
}

type camera interface {
	Position() (float64, float64)
	Rotation() float64
}

func (p *player) Position() (float64, float64) {
	return p.x, p.y
}

func (p *player) Rotation() float64 {
	return p.rotation
}

func (p *player) ToObjects() []Object {
	return []Object{
		{
			X:        playerX,
			Y:        playerY,
			Key:      "character_walking_" + p.key,
			Lifetime: p.lifetime,
		},
	}
}

const turnSpeed = 5.0
const moveSpeed = 50.0

const playerX = 200
const playerY = 200

// turning returns:
// -1 if player is turning left.
// 1 if player is turning right.
// 0 if player is not turning.
func (p *player) turning() int {
	return boolToInt[p.turnRight] - boolToInt[p.turnLeft]
}

// moving returns:
// 1 if player is moving forward.
// -1 if player is moving backwards.
// 0 if player is not moving.
func (p *player) moving() int {
	return boolToInt[p.moveForward] - boolToInt[p.moveBackward]
}

// isStanding is true if the player is neither moving nor turning.
func (p *player) isStanding() bool {
	return p.turning() == 0 && p.moving() == 0
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
