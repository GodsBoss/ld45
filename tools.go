package ld45

type toolID string

const (
	toolAxe     toolID = "axe"
	toolPickaxe toolID = "pickaxe"
	toolShovel  toolID = "shovel"
	toolSword   toolID = "sword"
)

type toolQuality int

const (
	toolNone toolQuality = iota
	toolWood
	toolStone
	toolIron
)
