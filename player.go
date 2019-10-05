package ld45

import "fmt"

type player struct {
	key string

	lifetime int

	health     intProperty
	saturation intProperty

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

	// inventory contains the player's gathered resources, e.g. wood, stone, etc.
	inventory map[itemID]int

	// equipment are the player's tools.
	equipment map[toolID]toolQuality
}

func (p *player) Position() (float64, float64) {
	return p.x, p.y
}

func (p *player) Rotation() float64 {
	return p.rotation
}

func (p *player) ToObjects() []Object {
	key := "character_walking_%s"
	if p.isStanding() {
		key = "character_standing_%s"
	}
	objects := []Object{
		{
			X:           playerX,
			Y:           playerY,
			Key:         fmt.Sprintf(key, p.key),
			Lifetime:    p.lifetime,
			GroundBound: true,
		},
	}
	for i := 0; i < p.health.current; i++ {
		objects = append(
			objects,
			Object{
				X:   2,
				Y:   2 + i*7,
				Key: "heart_full",
			},
		)
	}
	for i := p.health.current; i < maxHealth; i++ {
		objects = append(
			objects,
			Object{
				X:   2,
				Y:   2 + i*7,
				Key: "heart_empty",
			},
		)
	}
	for i := 0; i < p.saturation.current; i++ {
		objects = append(
			objects,
			Object{
				X:   10,
				Y:   2 + i*7,
				Key: "stomach_full",
			},
		)
	}
	for i := p.saturation.current; i < maxSaturation; i++ {
		objects = append(
			objects,
			Object{
				X:   10,
				Y:   2 + i*7,
				Key: "stomach_empty",
			},
		)
	}
	return objects
}

const turnSpeed = 5.0
const moveSpeed = 50.0

const playerX = 200
const playerY = 200

func inInteractionArea(x, y int) bool {
	return x >= playerX-5 && x <= playerX+5 && y <= playerY && y >= playerY-10
}

func inContact(x, y int) bool {
	return x >= playerX-playerContactSize && x <= playerX+playerContactSize && y >= playerY-playerContactSize && y <= playerY+playerContactSize
}

const playerContactSize = 3

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

func (p *player) filterInteractions(interactions []interaction) []interaction {
	result := make([]interaction, 0)
	for i := range interactions {
		if interactions[i].possible(p) {
			result = append(result, interactions[i])
		}
	}
	return result
}

var boolToInt = map[bool]int{
	false: 0,
	true:  1,
}

const maxHealth = 20
const maxSaturation = 20
