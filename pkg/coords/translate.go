package coords

// These constants can be used as factors for translation transformations to
// document the direction of the translation, e.g. Translation(Left * 3.5, Down * 2.5)
// transforms vectors by moving them 3.5 units to the left and 2.5 units downwards.
const (
	Left  = -1.0
	Right = 1.0
	Up    = -1.0
	Down  = 1.0
)

// Translation creates a transformation which translates vectors by (x, y).
func Translation(x, y float64) Transformation {
	return transformationMatrix{
		1.0, 0.0, x,
		0.0, 1.0, y,
	}
}

// TranslationByVector creates a transformation which translates vectors by the
// translation vector v.
func TranslationByVector(v Vector) Transformation {
	return Translation(v.X(), v.Y())
}
