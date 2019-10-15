package coords

// Vector is a 2D vector. Vectors are immutable. The zero value is the vector (0.0, 0.0).
type Vector struct {
	x float64
	y float64
}

// VectorFromCartesian creates a new 2D vector from cartesian coordinates x, y.
func VectorFromCartesian(x, y float64) Vector {
	return Vector{
		x: x,
		y: y,
	}
}

// X returns the vector's horizontal part.
func (v Vector) X() float64 {
	return v.x
}

// Y returns the vector's vertical part.
func (v Vector) Y() float64 {
	return v.y
}
