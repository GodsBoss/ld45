package coords

// Transformation converts vectors, i.e. takes input vectors and returns output vectors.
type Transformation interface {
	// Transform takes a 2D vector as input and returns an output vector. Two invocations
	// with the same vector as input MUST result in the same output vector (by value).
	Transform(input Vector) (output Vector)
}

// transformationMatrix is a 2D transformation matrix. If m is a matrix, as a
// mathematical representation it looks like this:
//
// m[0] m[1] m[2]
// m[3] m[4] m[5]
//    0    0    1
type transformationMatrix [6]float64

// Transform returns the product of this matrix and the input vector as output.
//
// The input vector v is represented like this:
//
// x
// y
// 1
//
// The resulting vector is the product of m * v.
func (matrix transformationMatrix) Transform(input Vector) Vector {
	return Vector{
		x: matrix[0]*input.x + matrix[1]*input.y + matrix[2],
		y: matrix[3]*input.x + matrix[4]*input.y + matrix[5],
	}
}
