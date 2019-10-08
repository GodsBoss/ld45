package ld45

import (
	"math"
)

type player struct {
	key string

	lifetime int

	health     intProperty
	saturation intProperty

	// conditionObjects stores the objects for health and saturation.
	conditionObjects []Object

	timeUntilSaturationLoss int

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

	remainingRegeneration intProperty

	// inventory contains the player's gathered resources, e.g. wood, stone, etc.
	inventory inventory

	// equipment are the player's tools.
	equipment map[toolID]toolQuality
}

func newPlayer(character string) *player {
	p := &player{
		key: character,
		health: intProperty{
			maximum: maxHealth,
			current: maxHealth,
		},
		saturation: intProperty{
			maximum: maxSaturation,
			current: maxSaturation,
		},
		timeUntilSaturationLoss: msPerSaturationLoss,
		remainingRegeneration: intProperty{
			maximum: regenerationDuration,
		},
		rotation: 0,
		x:        0,
		y:        0,
		inventory: inventory{
			possessions: make(map[itemID]int),
		},
		equipment: make(map[toolID]toolQuality),
	}
	p.conditionObjects = make([]Object, p.health.maximum+p.saturation.maximum)
	for i := 0; i < p.health.maximum; i++ {
		p.conditionObjects[i].X = 2
		p.conditionObjects[i].Y = 2 + i*7
	}
	for i := 0; i < p.saturation.maximum; i++ {
		p.conditionObjects[i+p.health.maximum].X = 10
		p.conditionObjects[i+p.health.maximum].Y = 2 + i*7
	}
	p.syncConditionObjects()
	return p
}

const msPerSaturationLoss = 30000

func (p *player) Position() (float64, float64) {
	return p.x, p.y
}

func (p *player) Rotation() float64 {
	return p.rotation
}

func (p *player) ToObjects() []Object {
	key := "character_walking_" + p.key
	if p.isStanding() {
		key = "character_standing_" + p.key
	}
	objects := []Object{
		{
			X:           playerX,
			Y:           playerY,
			Key:         key,
			Lifetime:    p.lifetime,
			GroundBound: true,
		},
	}
	objects = append(objects, p.conditionObjects...)
	objects = append(objects, p.inventory.Objects()...)
	return objects
}

func (p *player) syncConditionObjects() {
	for i := 0; i < p.health.current; i++ {
		p.conditionObjects[i].Key = "heart_full"
	}
	for i := p.health.current; i < p.health.maximum; i++ {
		p.conditionObjects[i].Key = "heart_empty"
	}
	for i := 0; i < p.saturation.current; i++ {
		p.conditionObjects[p.health.maximum+i].Key = "stomach_full"
	}
	for i := p.saturation.current; i < p.saturation.maximum; i++ {
		p.conditionObjects[p.health.maximum+i].Key = "stomach_empty"
	}
}

const turnSpeed = 5.0
const moveSpeed = 80.0

const playerX = 200
const playerY = 200

func inInteractionArea(x, y int) bool {
	return x >= playerX-5 && x <= playerX+5 && y <= playerY && y >= playerY-10
}

func distanceToPlayer(x, y int) float64 {
	return math.Sqrt(float64(x)*float64(x) + float64(y)*float64(y))
}

func inContact(x, y int) bool {
	return x >= playerX-playerContactSize && x <= playerX+playerContactSize && y >= playerY-playerContactSize && y <= playerY+playerContactSize
}

const playerContactSize = 5

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

func (p *player) Tick(ms int) {
	p.lifetime += ms
	p.remainingRegeneration.Dec(ms)
	if p.remainingRegeneration.IsMinimum() && !p.health.IsMaximum() && !p.saturation.IsMinimum() {
		p.health.Inc(1)
		p.saturation.Dec(1)
		p.remainingRegeneration.Inc(regenerationDuration)
		p.syncConditionObjects()
	}
	p.timeUntilSaturationLoss -= ms
	if p.timeUntilSaturationLoss <= 0 {
		p.timeUntilSaturationLoss += msPerSaturationLoss
		p.attemptSaturationLoss()
	}
}

func (p *player) attemptSaturationLoss() {
	if !p.saturation.IsMinimum() {
		p.saturation.Dec(1)
		p.syncConditionObjects()
		return
	}
	p.health.Dec(1)
	p.syncConditionObjects()
}

// regenerationDuration determines how much time (in ms) is needed to convert saturation into health.
const regenerationDuration = 2000
