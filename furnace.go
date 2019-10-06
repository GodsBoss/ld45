package ld45

type furnace struct {
	p *playing

	nopOnPlayerContact
	positionPartial

	burning           bool
	burningItem       itemID
	remainingBurnTime int
	lifetime          int
}

var burnTimes = map[itemID]int{
	itemIronOre: 20000,
	itemGoldOre: 30000,
}

func (furn *furnace) Interactions() []interaction {
	if furn.burning {
		return make([]interaction, 0)
	}
	return []interaction{
		newSimpleInteraction(
			"smelt_iron",
			possibleAll(
				minimalInventory(itemCoal, 1),
				minimalInventory(itemIronOre, 1),
			),
			func(_ int, p *playing) {
				furn.burning = true
				furn.burningItem = itemIronOre
				furn.remainingBurnTime = burnTimes[itemIronOre]
				p.player.inventory[itemCoal]--
				p.player.inventory[itemIronOre]--
			},
		),
		newSimpleInteraction(
			"smelt_gold",
			possibleAll(
				minimalInventory(itemCoal, 1),
				minimalInventory(itemGoldOre, 1),
			),
			func(_ int, p *playing) {
				furn.burning = true
				furn.burningItem = itemGoldOre
				furn.remainingBurnTime = burnTimes[itemGoldOre]
				p.player.inventory[itemCoal]--
				p.player.inventory[itemGoldOre]--
			},
		),
	}
}

var smeltingProducts = map[itemID]itemID{
	itemIronOre: itemIronIngot,
	itemGoldOre: itemGoldIngot,
}

func (furn *furnace) Tick(ms int) {
	furn.lifetime += ms
	if furn.burning {
		furn.remainingBurnTime -= ms
		if furn.remainingBurnTime <= 0 {
			furn.burning = false
			x, y := randomPositionAround(furn.x, furn.y, 5.0, 10.0)
			furn.p.interactibles.add(smeltingProducts[furn.burningItem].New(x, y))
		}
	}
}

func (furn *furnace) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, furn.x, furn.y)
	key := "furnace_off"
	if furn.burning {
		key = "furnace_burning"
	}
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         key,
			GroundBound: true,
			Lifetime:    furn.lifetime,
		},
	}
}
