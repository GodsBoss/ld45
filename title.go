package ld45

import (
	"github.com/GodsBoss/ld45/pkg/console"
)

type Title struct {
	faceLifetime int
}

func (title *Title) ID() string {
	return "title"
}

func (title *Title) Tick(ms int) {
	title.faceLifetime += ms
}

func (title *Title) Objects() []Object {
	return []Object{
		{
			X:        10,
			Y:        10,
			Key:      "face",
			Lifetime: title.faceLifetime,
		},
	}
}

func (title *Title) InvokeKeyEvent(event KeyEvent) {
	console.Log(event)
}
