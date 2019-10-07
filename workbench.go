package ld45

type workbench struct {
	nopOnPlayerContact
	positionPartial
	nopTick
	singleObject
}

func newWorkbench(x, y float64) *workbench {
	wb := &workbench{
		positionPartial: createPositionPartial(x, y),
		singleObject:    createSingleObject(x, y, true),
	}
	wb.singleObject.setKey("workbench")
	return wb
}

func (wb *workbench) ID() interactibleID {
	return "workbench"
}

func (wb *workbench) Interactions() []interaction {
	return recipeInteractions
}

type recipe struct {
	key       interactionID
	input     map[itemID]int
	output    func(*playing)
	condition func(*player) bool
}

func (r *recipe) toInteraction() interaction {
	return newSimpleInteraction(
		r.key,
		false,
		func(p *player) bool {
			if r.condition != nil && !r.condition(p) {
				return false
			}
			for inputItemID := range r.input {
				if !p.inventory.has(inputItemID, r.input[inputItemID]) {
					return false
				}
			}
			return true
		},
		func(_ int, p *playing) {
			for id := range r.input {
				p.player.inventory.add(id, -r.input[id])
			}
			r.output(p)
		},
	)
}

var recipeInteractions = func(recipes []recipe) []interaction {
	result := make([]interaction, len(recipes))
	for i := range recipes {
		result[i] = recipes[i].toInteraction()
	}
	return result
}(
	[]recipe{
		{
			key:       "interaction_axe_wood",
			input:     map[itemID]int{itemWood: 4},
			output:    workbenchToolOutput(toolAxe, toolWood),
			condition: maximumToolQuality(toolAxe, toolNone),
		},
		{
			key:       "interaction_pickaxe_wood",
			input:     map[itemID]int{itemWood: 4},
			output:    workbenchToolOutput(toolPickaxe, toolWood),
			condition: maximumToolQuality(toolPickaxe, toolNone),
		},
		{
			key:       "interaction_sword_wood",
			input:     map[itemID]int{itemWood: 4},
			output:    workbenchToolOutput(toolSword, toolWood),
			condition: maximumToolQuality(toolSword, toolNone),
		},
		{
			key: "interaction_axe_stone",
			input: map[itemID]int{
				itemWood: 1,
				itemRock: 3,
			},
			output:    workbenchToolOutput(toolAxe, toolStone),
			condition: maximumToolQuality(toolAxe, toolWood),
		},
		{
			key: "interaction_pickaxe_stone",
			input: map[itemID]int{
				itemWood: 1,
				itemRock: 3,
			},
			output:    workbenchToolOutput(toolPickaxe, toolStone),
			condition: maximumToolQuality(toolPickaxe, toolWood),
		},
		{
			key: "interaction_sword_stone",
			input: map[itemID]int{
				itemWood: 1,
				itemRock: 3,
			},
			output:    workbenchToolOutput(toolSword, toolStone),
			condition: maximumToolQuality(toolSword, toolWood),
		},
		{
			key: "interaction_axe_iron",
			input: map[itemID]int{
				itemWood:      1,
				itemIronIngot: 3,
			},
			output:    workbenchToolOutput(toolAxe, toolIron),
			condition: maximumToolQuality(toolAxe, toolStone),
		},
		{
			key: "interaction_pickaxe_iron",
			input: map[itemID]int{
				itemWood:      1,
				itemIronIngot: 3,
			},
			output:    workbenchToolOutput(toolPickaxe, toolIron),
			condition: maximumToolQuality(toolPickaxe, toolStone),
		},
		{
			key: "interaction_sword_iron",
			input: map[itemID]int{
				itemWood:      1,
				itemIronIngot: 3,
			},
			output:    workbenchToolOutput(toolSword, toolIron),
			condition: maximumToolQuality(toolSword, toolStone),
		},
		{
			key: "interaction_craft_furnace",
			input: map[itemID]int{
				itemRock: 4,
			},
			output: func(p *playing) {
				p.player.inventory.add(itemFurnace, 1)
			},
		},
		{
			key: "interaction_craft_crown",
			input: map[itemID]int{
				itemGoldIngot: 8,
				itemDiamond:   3,
			},
			output: func(p *playing) {
				p.result.SetVictory()
				p.transition("game_over")
			},
		},
	},
)

func workbenchToolOutput(id toolID, quality toolQuality) func(*playing) {
	return func(p *playing) {
		p.player.equipment[id] = quality
	}
}

func maximumToolQuality(id toolID, quality toolQuality) func(*player) bool {
	return func(p *player) bool {
		return p.equipment[id] <= quality
	}
}
