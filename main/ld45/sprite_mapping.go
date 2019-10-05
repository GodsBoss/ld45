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
	}
}
