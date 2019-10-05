package main

import (
	"github.com/GodsBoss/ld45/ui"
)

func createSpriteMapping() map[string]ui.Sprite {
	return map[string]ui.Sprite{
		"character_choice_1": ui.Sprite{
			X:               0,
			Y:               0,
			Width:           17,
			Height:          29,
			Frames:          2,
			FramesPerSecond: 4,
		},
		"character_choice_2": ui.Sprite{
			X:               0,
			Y:               29,
			Width:           17,
			Height:          29,
			Frames:          2,
			FramesPerSecond: 4,
		},
		"character_walking_1": ui.Sprite{
			X:               36,
			Y:               0,
			Width:           16,
			Height:          29,
			Frames:          8,
			FramesPerSecond: 12,
		},
		"character_standing_1": ui.Sprite{
			X:      68,
			Y:      0,
			Width:  16,
			Height: 29,
		},
		"character_walking_2": ui.Sprite{
			X:               36,
			Y:               29,
			Width:           16,
			Height:          29,
			Frames:          8,
			FramesPerSecond: 12,
		},
		"character_standing_2": ui.Sprite{
			X:      68,
			Y:      29,
			Width:  16,
			Height: 29,
		},
		"bush_3_berries": ui.Sprite{
			X:      200,
			Y:      11,
			Width:  19,
			Height: 13,
		},
		"bush_2_berries": ui.Sprite{
			X:      200,
			Y:      24,
			Width:  19,
			Height: 13,
		},
		"bush_1_berries": ui.Sprite{
			X:      200,
			Y:      37,
			Width:  19,
			Height: 13,
		},
		"bush_0_berries": ui.Sprite{
			X:      200,
			Y:      50,
			Width:  19,
			Height: 13,
		},
		"tree_1": ui.Sprite{
			X:      181,
			Y:      88,
			Width:  8,
			Height: 8,
		},
		"tree_2": ui.Sprite{
			X:      175,
			Y:      99,
			Width:  20,
			Height: 13,
		},
		"tree_3": ui.Sprite{
			X:      173,
			Y:      116,
			Width:  25,
			Height: 20,
		},
		"rock": ui.Sprite{
			X:      9,
			Y:      62,
			Width:  59,
			Height: 46,
		},
		"coal_ore": ui.Sprite{
			X:      68,
			Y:      62,
			Width:  59,
			Height: 46,
		},
		"iron_ore": ui.Sprite{
			X:      68,
			Y:      108,
			Width:  59,
			Height: 46,
		},
		"gold_ore": ui.Sprite{
			X:      9,
			Y:      108,
			Width:  59,
			Height: 46,
		},
		"item_berry": ui.Sprite{
			X:      166,
			Y:      3,
			Width:  3,
			Height: 3,
		},
		"item_wood": ui.Sprite{
			X:      167,
			Y:      9,
			Width:  5,
			Height: 5,
		},
		"item_sapling": ui.Sprite{
			X:      167,
			Y:      17,
			Width:  6,
			Height: 3,
		},
		"item_rock": ui.Sprite{
			X:      181,
			Y:      6,
			Width:  4,
			Height: 3,
		},
		"item_coal": ui.Sprite{
			X:      181,
			Y:      12,
			Width:  4,
			Height: 3,
		},
		"item_iron_ore": ui.Sprite{
			X:      181,
			Y:      18,
			Width:  4,
			Height: 3,
		},
		"item_gold_ore": ui.Sprite{
			X:      181,
			Y:      24,
			Width:  4,
			Height: 3,
		},
		"heart_full": ui.Sprite{
			X:      175,
			Y:      36,
			Width:  7,
			Height: 6,
		},
		"heart_empty": ui.Sprite{
			X:      183,
			Y:      36,
			Width:  7,
			Height: 6,
		},
		"stomach_full": ui.Sprite{
			X:      175,
			Y:      43,
			Width:  7,
			Height: 6,
		},
		"stomach_empty": ui.Sprite{
			X:      183,
			Y:      43,
			Width:  7,
			Height: 6,
		},
	}
}
