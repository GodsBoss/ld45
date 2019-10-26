package coords

// Scale scales coordinates horizontally and vertically.
func Scale(horizontal, vertical float64) Transformation {
	return transformationMatrix{
		horizontal, 0.0, 0.0,
		0.0, vertical, 0.0,
	}
}
