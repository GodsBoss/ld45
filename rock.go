package ld45

import "math/rand"

type rock struct {
	positionPartial
	nopOnPlayerContact
	storedInteractions

	key                 rockType
	remainingDurability int
}

func newRock(x, y float64, key rockType) *rock {
	r := &rock{
		positionPartial:     createPositionPartial(x, y),
		key:                 key,
		remainingDurability: key.durability(),
	}
	r.interactions = []interaction{
		newSimpleInteraction(
			"interaction_harvest_rock",
			true,
			func(p *player) bool {
				return p.equipment[toolPickaxe] >= r.key.minimumPickaxeLevel()
			},
			func(index int, p *playing) {
				r.remainingDurability -= int(p.player.equipment[toolPickaxe]) * 2
				if r.remainingDurability <= 0 {
					rewards := r.key.reward()
					for iID := range rewards {
						for i := 0; i < rewards[iID]; i++ {
							x, y := randomPositionAround(r.x, r.y, 20.0, 30.0)
							p.interactibles.add(
								iID.New(x, y),
							)
						}
					}
					p.interactibles.remove(index)
				}
			},
		),
	}
	return r
}

type rockType string

const (
	rockStone   rockType = "rock_stone"
	rockCoal    rockType = "rock_coal"
	rockIronOre rockType = "rock_iron_ore"
	rockGoldOre rockType = "rock_gold_ore"
	rockDiamond rockType = "rock_diamond"
)

func (t rockType) minimumPickaxeLevel() toolQuality {
	q, ok := minimumPickaxeLevels[t]
	if ok {
		return q
	}
	return pickaxeLevelForUnknownRocks
}

const pickaxeLevelForUnknownRocks = toolWood

var minimumPickaxeLevels = map[rockType]toolQuality{
	rockStone:   toolWood,
	rockCoal:    toolStone,
	rockIronOre: toolStone,
	rockGoldOre: toolIron,
	rockDiamond: toolIron,
}

func (t rockType) durability() int {
	dur, ok := durabilitiesPerRockType[t]
	if ok {
		return dur
	}
	return durabilityForUnknownRocks
}

const durabilityForUnknownRocks = 5

var durabilitiesPerRockType = map[rockType]int{
	rockStone:   10,
	rockCoal:    10,
	rockIronOre: 15,
	rockGoldOre: 15,
	rockDiamond: 20,
}

func (t rockType) reward() map[itemID]int {
	cfg := rewardsPerRockType[t]
	reward := make(map[itemID]int)
	for id := range cfg {
		bonus := 0
		if cfg[id].max > cfg[id].min {
			bonus = rand.Intn(cfg[id].max - cfg[id].min)
		}
		reward[id] = cfg[id].min + bonus
	}
	return reward
}

var rewardsPerRockType = map[rockType]map[itemID]struct {
	min int
	max int
}{
	rockStone: {
		itemRock: {min: 2, max: 4},
	},
	rockCoal: {
		itemRock: {min: 1, max: 2},
		itemCoal: {min: 2, max: 3},
	},
	rockIronOre: {
		itemRock:    {min: 1, max: 2},
		itemIronOre: {min: 1, max: 2},
	},
	rockGoldOre: {
		itemRock:    {min: 1, max: 2},
		itemGoldOre: {min: 1, max: 2},
	},
	rockDiamond: {
		itemRock:    {min: 1, max: 2},
		itemDiamond: {min: 1, max: 1},
	},
}

func (r *rock) ID() interactibleID {
	return interactibleID(r.key)
}

func (r *rock) Tick(ms int) {}

func (r *rock) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, r.x, r.y)
	return []Object{
		{
			X:           x,
			Y:           y,
			Key:         string(r.key),
			Lifetime:    0,
			GroundBound: true,
		},
	}
}
