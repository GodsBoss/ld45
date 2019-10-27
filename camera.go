package ld45

import (
	"github.com/GodsBoss/ld45/pkg/coords"
)

type Camera struct {
	position coords.Vector
	rotation float64
}

func (camera *Camera) Position() coords.Vector {
	return camera.position
}

func (camera *Camera) MoveTo(position coords.Vector) {
	camera.position = position
}

func (camera *Camera) MoveBy(positionDelta coords.Vector) {
	camera.position = coords.TranslationByVector(positionDelta).Transform(camera.position)
}

func (camera *Camera) Rotation() float64 {
	return camera.rotation
}

func (camera *Camera) RotateTo(rotation float64) {
	camera.rotation = rotation
}

func (camera *Camera) RotateBy(rotationDelta float64) {
	camera.RotateTo(camera.Rotation() + rotationDelta)
}

// ScreenPosition calculates the screen position for another object.
func (camera *Camera) ScreenPosition(position coords.Vector) coords.Vector {
	position = coords.Combine(
		coords.Rotation(-camera.rotation),
		coords.TranslationByVector(coords.Scale(-1, -1).Transform(camera.position)),
	).Transform(position)
	return screenOffsetTransform.Transform(position)
}

var screenOffsetTransform = coords.TranslationByVector(coords.VectorFromCartesian(200, 150))
