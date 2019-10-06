package ld45

import (
	"math"
	"sort"
)

type playing struct {
	choice *characterChoice

	player         *player
	interactibles  *interactibles
	helpEnabled    bool
	interactionHub *interactionHub
}

func (playing *playing) ID() string {
	return "playing"
}

func (playing *playing) Init() {
	playing.interactibles = newInteractibles()
	playing.player = newPlayer(playing.choice.Get())
	playing.interactionHub = &interactionHub{
		playing:           playing,
		chosenInteraction: make(map[interactibleID]interactionID),
	}
	predefinedInteractibles := []interactible{
		&bush{
			positionPartial: createPositionPartial(0.0, -50.0),
			growth: intProperty{
				maximum: 3,
				current: 1,
			},
		},
		&bush{
			positionPartial: createPositionPartial(-25.0, -60.0),
			growth: intProperty{
				maximum: 3,
				current: 2,
			},
		},
		&bush{
			positionPartial: createPositionPartial(60.0, -25.0),
			growth: intProperty{
				maximum: 3,
				current: 3,
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
		itemIronIngot.New(-10.0, -100.0),
		itemGoldIngot.New(10.0, -100.0),
		&workbench{
			positionPartial: createPositionPartial(-75.0, -5.0),
		},
		&furnace{
			p:               playing,
			positionPartial: createPositionPartial(5.0, -80.0),
		},
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
	playing.interactionHub.Tick(ms)
}

func (playing *playing) Objects() []Object {
	objects := make(Objects, 0)
	objects = append(objects, playing.player.ToObjects()...)
	playing.interactibles.each(func(_ int, i interactible) {
		objects = append(objects, i.ToObjects(playing.player)...)
	})
	objects = append(objects, playing.interactionHub.Objects()...)
	sort.Sort(objects)
	return objects
}

func (playing *playing) triggerHelp() {
	playing.helpEnabled = !playing.helpEnabled
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
	case "i":
		if event.Type == KeyDown {
			playing.interactionHub.playerInteractsDirectly()
		}
	case "k":
		if event.Type == KeyDown {
			playing.interactionHub.playerInteractsIndirectly()
		}
	case "j":
		if event.Type == KeyDown {
			playing.interactionHub.changeIndirectPlayerChoice(-1)
		}
	case "l":
		if event.Type == KeyDown {
			playing.interactionHub.changeIndirectPlayerChoice(1)
		}
	case "h":
		playing.triggerHelp()
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
