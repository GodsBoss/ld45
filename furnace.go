package ld45

type furnace struct {
	p *playing

	nopOnPlayerContact
	positionPartial

	smelting             *smelting
	smeltingInteractions []interaction

	remainingBurnTime int
	lifetime          int
}

func newFurnace(p *playing, x, y float64) *furnace {
	return &furnace{
		p:                    p,
		positionPartial:      createPositionPartial(x, y),
		smeltingInteractions: interactionsFromSmeltings(smeltings),
	}
}

func (furn *furnace) isSmelting() bool {
	return furn.smelting != nil
}

func (furn *furnace) ID() interactibleID {
	return "furnace"
}

func (furn *furnace) Interactions() []interaction {
	if furn.isSmelting() {
		return make([]interaction, 0)
	}
	return furn.smeltingInteractions
}

type smelting struct {
	id         interactionID
	burntimeMS int
	input      map[itemID]int
	output     map[itemID]int
}

func (sm *smelting) toInteraction() interaction {
	return newSimpleInteraction(
		sm.id,
		false,
		func(p *player) bool {
			for costItemID := range sm.input {
				if !p.inventory.has(costItemID, sm.input[costItemID]) {
					return false
				}
			}
			return true
		},
		func(id int, p *playing) {
			furnace, ok := p.interactibles.m[id].(*furnace)

			// Should not happen, but better guard against this!
			if !ok {
				log("furnace interaction without a furnace")
				return
			}

			furnace.smelting = sm
			furnace.remainingBurnTime = sm.burntimeMS

			for iID := range sm.input {
				p.player.inventory.add(iID, -sm.input[iID])
			}
		},
	)
}

var smeltings = []*smelting{
	&smelting{
		id:         "interaction_smelt_iron",
		burntimeMS: 20000,
		input: map[itemID]int{
			itemCoal:    1,
			itemIronOre: 1,
		},
		output: map[itemID]int{
			itemIronIngot: 1,
		},
	},
	&smelting{
		id:         "interaction_smelt_gold",
		burntimeMS: 30000,
		input: map[itemID]int{
			itemCoal:    1,
			itemGoldOre: 1,
		},
		output: map[itemID]int{
			itemGoldIngot: 1,
		},
	},
	&smelting{
		id:         "interaction_smelt_coal",
		burntimeMS: 5000,
		input: map[itemID]int{
			itemWood: 3,
		},
		output: map[itemID]int{
			itemCoal: 1,
		},
	},
}

func interactionsFromSmeltings(smeltings []*smelting) []interaction {
	result := make([]interaction, len(smeltings))
	for i := range smeltings {
		result[i] = smeltings[i].toInteraction()
	}
	return result
}

func (furn *furnace) Tick(ms int) {
	furn.lifetime += ms
	if furn.isSmelting() {
		furn.remainingBurnTime -= ms
		if furn.remainingBurnTime <= 0 {
			for iID := range furn.smelting.output {
				for i := 0; i < furn.smelting.output[iID]; i++ {
					x, y := randomPositionAround(furn.x, furn.y, 5.0, 10.0)
					furn.p.interactibles.add(iID.New(x, y))
				}
			}
			furn.smelting = nil
		}
	}
}

func (furn *furnace) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, furn.x, furn.y)
	key := "furnace_off"
	if furn.isSmelting() {
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
