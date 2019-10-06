package ld45

import (
	"math"
	"math/rand"
)

type itemID string

func (id itemID) New(x, y float64) *item {
	return &item{
		id:              id,
		positionPartial: createPositionPartial(x, y),
	}
}

const (
	itemBerry     itemID = "item_berry"
	itemWood      itemID = "item_wood"
	itemSapling   itemID = "item_sapling"
	itemRock      itemID = "item_rock"
	itemCoal      itemID = "item_coal"
	itemIronOre   itemID = "item_iron_ore"
	itemGoldOre   itemID = "item_gold_ore"
	itemIronIngot itemID = "item_iron_ingot"
	itemGoldIngot itemID = "item_gold_ingot"
	itemFurnace   itemID = "item_furnace"
)

type item struct {
	id itemID

	noInteractions
	nopTick
	positionPartial
}

func (i *item) ID() interactibleID {
	return interactibleID(string(i.id))
}

func (i *item) OnPlayerContact(id int, p *playing) {
	p.player.inventory.add(i.id, 1)
	p.interactibles.remove(id)
}

func (i *item) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, i.x, i.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         string(i.ID()),
			Lifetime:    0,
			GroundBound: true,
		},
	}
}

func randomPositionAround(x, y, minR, maxR float64) (float64, float64) {
	angle := rand.Float64() * math.Pi * 2
	d := rand.Float64()*(maxR-minR) + minR
	return x + math.Sin(angle)*d, y + math.Cos(angle)*d
}
