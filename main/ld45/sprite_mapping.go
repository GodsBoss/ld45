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
	}
}
