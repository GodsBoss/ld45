package ld45

type toolID string

const (
	toolAxe     toolID = "axe"
	toolPickaxe        = "pickaxe"
	toolShovel         = "shovel"
	toolSword          = "sword"
)

type toolQuality int

const (
	toolNone toolQuality = iota
	toolWood
	toolStone
	toolIron
)
