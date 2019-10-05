package ui

import (
	"github.com/GodsBoss/ld45"
	"github.com/GodsBoss/ld45/pkg/dom"
)

type Renderer struct {
	Zoom              int
	Ctx               *dom.Context2D
	ImageSource       dom.ImageSource
	SpriteMapping     map[string]Sprite
	BackgroundMapping map[string]dom.FillStyle
}

var defaultFillStyle = dom.Color("#000000")

func (renderer *Renderer) Draw(game *ld45.Game) {
	bg, ok := renderer.BackgroundMapping[game.StateID()]
	if !ok {
		bg = defaultFillStyle
	}
	renderer.Ctx.FillStyle(bg)
	renderer.Ctx.FillRect(0, 0, float64(renderer.Ctx.Width()), float64(renderer.Ctx.Height()))
	for _, object := range game.Objects() {
		renderer.drawObject(object)
	}
}

func (renderer *Renderer) drawObject(object ld45.Object) {
	zoom := renderer.Zoom
	if zoom <= 0 {
		zoom = 1
	}
	sprite, ok := renderer.SpriteMapping[object.Key]
	if !ok {
		return
	}
	frame := 0
	if sprite.Frames > 1 {
		frame = ((object.Lifetime * sprite.FramesPerSecond) / 1000) % sprite.Frames
	}
	destinationX, destinationY := object.X, object.Y
	if object.GroundBound {
		destinationY -= sprite.Height
		destinationX -= sprite.Width / 2
	}
	renderer.Ctx.DrawImage(
		renderer.ImageSource,
		sprite.X+(sprite.Width*frame),
		sprite.Y,
		sprite.Width,
		sprite.Height,
		destinationX*zoom,
		destinationY*zoom,
		sprite.Width*zoom,
		sprite.Height*zoom,
	)
}

type Sprite struct {
	// X and Y are the coordinates of the first sprite within the source image.
	X int
	Y int

	// Width and Height are the size of the sprites.
	Width  int
	Height int

	// Frames is the number of frames this sprite animation has, if any. If 0,
	// 1 is assumed.
	Frames int

	// FramesPerSecond is the animation speed.
	FramesPerSecond int
}
