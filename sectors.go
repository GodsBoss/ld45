package ld45

import (
	"math"
	"math/rand"
)

// sectors is responsible for generating objects in neighbouring sectors when
// the player enters a specific sector.
type sectors struct {
	alreadyGenerated map[sectorID]struct{}
	sectorWidth      float64
	sectorHeight     float64
	generate         func(id sectorID, s sector)
}

func newSectors(sectorWidth, sectorHeight float64, generate func(id sectorID, s sector)) *sectors {
	return &sectors{
		sectorWidth:      sectorWidth,
		sectorHeight:     sectorHeight,
		generate:         generate,
		alreadyGenerated: make(map[sectorID]struct{}),
	}
}

func (s *sectors) playerMovesTo(x, y float64) {
	currentSector := s.positionToSectorID(x, y)
	if _, ok := s.alreadyGenerated[currentSector]; ok {
		return
	}
	candidates := currentSector.sectorIncludingNeighbours()
	for i := range candidates {
		s.attemptGenerate(candidates[i])
	}
}

func (s *sectors) positionToSectorID(x, y float64) sectorID {
	return sectorID{
		X: int(math.Floor(x/s.sectorWidth + 0.5)),
		Y: int(math.Floor(y/s.sectorHeight + 0.5)),
	}
}

func (s *sectors) attemptGenerate(id sectorID) {
	if _, ok := s.alreadyGenerated[id]; ok {
		return
	}
	cx, cy := float64(id.X)*s.sectorWidth, float64(id.Y)*s.sectorHeight
	s.generate(
		id,
		sector{
			Left:   cx - s.sectorWidth*0.5,
			Top:    cy - s.sectorHeight*0.5,
			Right:  cx + s.sectorWidth*0.5,
			Bottom: cy + s.sectorHeight*0.5,
		},
	)
	s.alreadyGenerated[id] = struct{}{}
}

type sectorID struct {
	X int
	Y int
}

func (id sectorID) sectorIncludingNeighbours() []sectorID {
	ids := make([]sectorID, 9)
	index := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			ids[index] = sectorID{
				X: id.X + dx,
				Y: id.Y + dy,
			}
			index++
		}
	}
	return ids
}

type sector struct {
	Left   float64
	Right  float64
	Bottom float64
	Top    float64
}

// Random returns a random point within the sector.
func (s sector) Random() (float64, float64) {
	return rand.Float64()*(s.Right-s.Left) + s.Left, rand.Float64()*(s.Bottom-s.Top) + s.Top
}
