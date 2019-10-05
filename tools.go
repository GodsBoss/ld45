package ld45

type toolID string

type toolQuality int

const (
	toolNone toolQuality = iota
	toolWood
	toolStone
	toolIron
)
