package ld45

type interactible interface {
	ID() interactibleID
	Interactions() []interaction
	OnPlayerContact(id int, p *playing)
	Position() (float64, float64)
	Tick(ms int)
	ToObjects(Camera) []Object
}

type interactibleID string

var noInteractions = make([]interaction, 0)

type storedInteractions struct {
	interactions []interaction
}

func (si storedInteractions) Interactions() []interaction {
	return si.interactions
}

type nopOnPlayerContact struct{}

func (contact nopOnPlayerContact) OnPlayerContact(_ int, _ *playing) {}

type positionPartial struct {
	x float64
	y float64
}

type nopTick struct{}

func (tick nopTick) Tick(ms int) {}

func createPositionPartial(x, y float64) positionPartial {
	return positionPartial{
		x: x,
		y: y,
	}
}

func (p positionPartial) Position() (float64, float64) {
	return p.x, p.y
}

type interactibles struct {
	sectors *sectors

	lastID                 int
	idsToSector            map[int]sectorID
	interactiblesPerSector map[sectorID]map[int]interactible
}

func newInteractibles(secs *sectors) *interactibles {
	return &interactibles{
		sectors:                secs,
		idsToSector:            make(map[int]sectorID),
		interactiblesPerSector: make(map[sectorID]map[int]interactible),
	}
}

func (is *interactibles) get(id int) interactible {
	sID, ok := is.idsToSector[id]
	if !ok {
		return nil
	}
	return is.interactiblesPerSector[sID][id]
}

func (is *interactibles) add(i interactible) int {
	is.lastID++
	x, y := i.Position()
	sID := is.sectors.positionToSectorID(x, y)
	is.idsToSector[is.lastID] = sID
	if _, ok := is.interactiblesPerSector[sID]; !ok {
		is.interactiblesPerSector[sID] = make(map[int]interactible)
	}
	is.interactiblesPerSector[sID][is.lastID] = i
	return is.lastID
}

func (is *interactibles) remove(id int) {
	sID := is.idsToSector[id]
	delete(is.interactiblesPerSector[sID], id)
	delete(is.idsToSector, id)
}

func (is *interactibles) each(f func(id int, i interactible)) {
	for sID := range is.interactiblesPerSector {
		for id := range is.interactiblesPerSector[sID] {
			f(id, is.interactiblesPerSector[sID][id])
		}
	}
}

func (is *interactibles) eachWithin(sIDs []sectorID, f func(id int, i interactible)) {
	for i := range sIDs {
		for id := range is.interactiblesPerSector[sIDs[i]] {
			f(id, is.interactiblesPerSector[sIDs[i]][id])
		}
	}
}
