package ld45

type itemID string

func (id itemID) New(x, y float64) *item {
	return &item{
		ID:              id,
		positionPartial: createPositionPartial(x, y),
	}
}

const (
	itemBerry   itemID = "item_berry"
	itemWood    itemID = "item_wood"
	itemSapling itemID = "item_sapling"
	itemRock    itemID = "item_rock"
	itemCoal    itemID = "item_coal"
	itemIronOre itemID = "item_iron_ore"
	itemGoldOre itemID = "item_gold_ore"
)

type item struct {
	ID itemID

	noInteractions
	nopTick
	positionPartial
}

func (i *item) OnPlayerContact(id int, p *playing) {
	p.player.inventory[i.ID]++
	p.interactibles.remove(id)
}

func (i *item) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, i.x, i.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         string(i.ID),
			Lifetime:    0,
			GroundBound: true,
		},
	}
}
