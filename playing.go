package ld45

import (
	"github.com/GodsBoss/ld45/pkg/coords"

	"math"
	"math/rand"
	"sort"
)

type playing struct {
	transitioner
	choice *characterChoice
	result *playResult

	player         *player
	interactibles  *interactibles
	helpEnabled    bool
	interactionHub *interactionHub
	sectors        *sectors
}

func (playing *playing) ID() string {
	return "playing"
}

func (playing *playing) Init() {
	playing.sectors = newSectors(400.0, 300.0, playing.generateSector)
	playing.interactibles = newInteractibles(playing.sectors)
	playing.player = newPlayer(playing.choice.Get())
	playing.interactionHub = &interactionHub{
		playing:           playing,
		chosenInteraction: make(map[interactibleID]interactionID),
		defaultActions: &playerActions{
			storedInteractions: storedInteractions{
				interactions: playerActionsInteractions(),
			},
		},
	}
	predefinedInteractibles := []interactible{
		newBush(-40.0, -60.0, 2),
		newBush(50.0, -50.0, 2),
		newTree(10.0, 35.0, 3),
		newTree(-75.0, -30.0, 3),
		newTree(-90.0, -60.0, 3),
		newTree(90.0, -70.0, 2),
		newTree(-180.0, 100.0, 3),
		newRock(170.0, 95.0, rockStone),
	}
	for i := range predefinedInteractibles {
		playing.interactibles.add(predefinedInteractibles[i])
	}
}

func (playing *playing) generateSector(id sectorID, s sector) {
	// Start sector is special, don't generate anything here.
	if id.X == 0 && id.Y == 0 {
		return
	}
	bx, by := s.Random()
	bg := rand.Intn(3)
	playing.interactibles.add(
		newBush(bx, by, bg),
	)
	for i := 0; i < 5; i++ {
		treeSize := 1 + rand.Intn(2) + rand.Intn(2)
		tx, ty := s.Random()
		playing.interactibles.add(
			newTree(tx, ty, treeSize),
		)
	}
	for i := 0; i < 4; i++ {
		rx, ry := s.Random()
		playing.interactibles.add(
			newRock(rx, ry, rockStone),
		)
	}
	if id.X > 3 || id.X < -3 || id.Y > 3 || id.Y < -3 {
		if rand.Float64() < 0.75 {
			playing.generateSectorMedium(id, s)
		} else {
			playing.generateSectorRich(id, s)
		}
	}
}

func (playing *playing) generateSectorMedium(id sectorID, s sector) {
	coalRocks := 1 + rand.Intn(2)
	for i := 0; i < coalRocks; i++ {
		rx, ry := s.Random()
		playing.interactibles.add(
			newRock(rx, ry, rockCoal),
		)
	}
	iox, ioy := s.Random()
	playing.interactibles.add(
		newRock(iox, ioy, rockIronOre),
	)
	if rand.Intn(2)*rand.Intn(2) == 1 {
		gx, gy := s.Random()
		playing.interactibles.add(
			newRock(gx, gy, rockGoldOre),
		)
	}
}

func (playing *playing) generateSectorRich(id sectorID, s sector) {
	coalRocks := 1 + rand.Intn(2)
	for i := 0; i < coalRocks; i++ {
		rx, ry := s.Random()
		playing.interactibles.add(
			newRock(rx, ry, rockCoal),
		)
	}
	gx, gy := s.Random()
	playing.interactibles.add(
		newRock(gx, gy, rockGoldOre),
	)
	if rand.Intn(2) == 1 {
		dx, dy := s.Random()
		playing.interactibles.add(
			newRock(dx, dy, rockDiamond),
		)
	}
}

func (playing *playing) Tick(ms int) {
	playing.player.Tick(ms)
	timeFactor := float64(ms) / 1000
	playing.player.rotation += turnSpeed * float64(playing.player.turn.asInt()) * timeFactor

	movement := coords.VectorFromCartesian(0, -1)
	movement = coords.Scale(0, moveSpeed*timeFactor).Transform(movement)
	movement = coords.Rotation(playing.player.rotation).Transform(movement)
	movement = coords.Rotation(moveStrafeAngles[moveStrafeAngleKey{move: playing.player.move.asInt(), strafe: playing.player.strafe.asInt()}]).Transform(movement)
	if playing.player.move.isSome() || playing.player.strafe.isSome() {
		playing.player.x += movement.X()
		playing.player.y += movement.Y()
	}
	playing.interactibles.each(func(id int, i interactible) {
		i.Tick(ms)
		ix, iy := i.Position()
		x, y := calculateScreenPosition(playing.player, ix, iy)
		if inContact(x, y) {
			i.OnPlayerContact(id, playing)
		}
	})
	playing.interactionHub.Tick(ms)
	playing.sectors.playerMovesTo(playing.player.x, playing.player.y)
	if playing.player.health.IsMinimum() {
		playing.result.SetDefeat()
		playing.transition("game_over")
	}
}

type moveStrafeAngleKey struct {
	move   int
	strafe int
}

var moveStrafeAngles = map[moveStrafeAngleKey]float64{
	{move: 1, strafe: 0}:   coords.FullAngle * 0.0,
	{move: 1, strafe: 1}:   coords.FullAngle * 0.125,
	{move: 0, strafe: 1}:   coords.FullAngle * 0.25,
	{move: -1, strafe: 1}:  coords.FullAngle * 0.375,
	{move: -1, strafe: 0}:  coords.FullAngle * 0.5,
	{move: -1, strafe: -1}: coords.FullAngle * 0.625,
	{move: 0, strafe: -1}:  coords.FullAngle * 0.75,
	{move: 1, strafe: -1}:  coords.FullAngle * 0.875,
}

func (playing *playing) Objects() []Object {
	objects := make(Objects, 0)
	objects = append(objects, playing.player.ToObjects()...)
	px, py := playing.player.Position()
	playing.interactibles.eachWithin(
		playing.sectors.positionToSectorID(px, py).sectorIncludingNeighbours(),
		func(_ int, i interactible) {
			objects = append(objects, i.ToObjects(playing.player)...)
		},
	)
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
			playing.player.turn.enableSecond()
		}
		if event.Type == KeyUp {
			playing.player.turn.disableSecond()
		}
	case "d":
		if event.Type == KeyDown {
			playing.player.turn.enableFirst()
		}
		if event.Type == KeyUp {
			playing.player.turn.disableFirst()
		}
	case "w":
		if event.Type == KeyDown {
			playing.player.move.enableFirst()
		}
		if event.Type == KeyUp {
			playing.player.move.disableFirst()
		}
	case "s":
		if event.Type == KeyDown {
			playing.player.move.enableSecond()
		}
		if event.Type == KeyUp {
			playing.player.move.disableSecond()
		}
	case "q":
		if event.Type == KeyDown {
			playing.player.strafe.enableSecond()
		}
		if event.Type == KeyUp {
			playing.player.strafe.disableSecond()
		}
	case "e":
		if event.Type == KeyDown {
			playing.player.strafe.enableFirst()
		}
		if event.Type == KeyUp {
			playing.player.strafe.disableFirst()
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

func relativePosition(ox, oy, dx, dy, rotation float64) (float64, float64) {
	return ox - dx*math.Cos(rotation) - dy*math.Sin(rotation), oy - dx*math.Sin(rotation) + dy*math.Cos(rotation)
}
