package ld45

type furnace struct {
	nopOnPlayerContact
	positionPartial

	burning     bool
	burningItem itemID
	lifetime    int
}

func (furn *furnace) Interactions() []interaction {
	if furn.burning {
		return make([]interaction, 0)
	}
	return []interaction{
		newSimpleInteraction(
			possibleAll(
				minimalInventory(itemCoal, 1),
				minimalInventory(itemIronOre, 1),
			),
			func(_ int, p *playing) {
				furn.burning = true
				furn.burningItem = itemIronOre
				p.player.inventory[itemCoal]--
				p.player.inventory[itemIronOre]--
			},
		),
		newSimpleInteraction(
			possibleAll(
				minimalInventory(itemCoal, 1),
				minimalInventory(itemGoldOre, 1),
			),
			func(_ int, p *playing) {
				furn.burning = true
				furn.burningItem = itemGoldOre
				p.player.inventory[itemCoal]--
				p.player.inventory[itemGoldOre]--
			},
		),
	}
}

func (furn *furnace) Tick(ms int) {
	furn.lifetime += ms
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
