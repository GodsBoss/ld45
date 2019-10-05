package ld45

import (
	"math"
	"math/rand"
	"sort"
)

type playing struct {
	choice *characterChoice

	player        *player
	interactibles *interactibles
}

func (playing *playing) ID() string {
	return "playing"
}

func (playing *playing) Init() {
	playing.interactibles = newInteractibles()
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
		rotation:  0,
		x:         0,
		y:         0,
		inventory: make(map[itemID]int),
		equipment: make(map[toolID]toolQuality),
	}
	predefinedInteractibles := []interactible{
		&bush{
			positionPartial: createPositionPartial(0.0, -50.0),
			growth: intProperty{
				maximum: 3,
			},
		},
		&bush{
			positionPartial: createPositionPartial(-25.0, -60.0),
			growth: intProperty{
				maximum: 3,
			},
		},
		&bush{
			positionPartial: createPositionPartial(60.0, -25.0),
			growth: intProperty{
				maximum: 3,
			},
		},
		newTree(-80.0, -80.0, 1),
		newTree(-120.0, 25.0, 2),
		newTree(90.0, 5.0, 3),
		&rock{
			positionPartial: createPositionPartial(100.0, 25.0),
			key:             "rock",
		},
		&rock{
			positionPartial: createPositionPartial(-90.0, -75.0),
			key:             "coal_ore",
		},
		&rock{
			positionPartial: createPositionPartial(55.0, 40.0),
			key:             "iron_ore",
		},
		&rock{
			positionPartial: createPositionPartial(60.0, -90.0),
			key:             "gold_ore",
		},
		itemBerry.New(-30.0, 40.0),
		itemWood.New(-15.0, 50.0),
		itemSapling.New(0.0, 35.0),
		itemCoal.New(15.0, 40.0),
		itemIronOre.New(30.0, 35.0),
		itemRock.New(-45.0, 30.0),
		itemGoldOre.New(45.0, 40.0),
	}
	for i := range predefinedInteractibles {
		playing.interactibles.add(predefinedInteractibles[i])
	}
}

func (playing *playing) Tick(ms int) {
	playing.player.lifetime += ms
	playing.player.rotation += turnSpeed * float64(playing.player.turning()) * float64(ms) / 1000
	playing.player.x += float64(playing.player.moving()) * moveSpeed * math.Sin(playing.player.rotation) * float64(ms) / 1000
	playing.player.y += float64(playing.player.moving()) * moveSpeed * -math.Cos(playing.player.rotation) * float64(ms) / 1000
	playing.interactibles.each(func(id int, i interactible) {
		i.Tick(ms)
		ix, iy := i.Position()
		x, y := calculateScreenPosition(playing.player, ix, iy)
		if inContact(x, y) {
			i.OnPlayerContact(id, playing)
		}
	})
}

func (playing *playing) Objects() []Object {
	objects := make(Objects, 0)
	objects = append(objects, playing.player.ToObjects()...)
	playing.interactibles.each(func(_ int, i interactible) {
		objects = append(objects, i.ToObjects(playing.player)...)
	})
	sort.Sort(objects)
	return objects
}

func (playing *playing) playerInteracts() {
	interactionCandidates := make([]func(), 0)
	playing.interactibles.each(func(id int, i interactible) {
		ix, iy := i.Position()
		x, y := calculateScreenPosition(playing.player, ix, iy)
		if inInteractionArea(x, y) {
			interactibleInteractions := playing.player.filterInteractions(i.Interactions())
			for j := range interactibleInteractions {
				interactionCandidates = append(
					interactionCandidates,
					func(id int, candidate interaction) func() {
						return func() {
							candidate.invoke(id, playing)
						}
					}(id, interactibleInteractions[j]),
				)
			}
		}
	})
	if len(interactionCandidates) == 0 {
		return
	}
	interactionCandidates[rand.Intn(len(interactionCandidates))]()
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

func calculateScreenPosition(cam camera, ox, oy float64) (x int, y int) {
	cx, cy := cam.Position()
	dx, dy := cx-ox, cy-oy
	rx := -dx*math.Cos(cam.Rotation()) - dy*math.Sin(cam.Rotation())
	ry := dx*math.Sin(cam.Rotation()) - dy*math.Cos(cam.Rotation())
	return int(rx) + playerX, int(ry) + playerY
}
