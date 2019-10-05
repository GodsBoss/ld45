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
	}
}
