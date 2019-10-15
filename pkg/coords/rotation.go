package coords

import (
	"math"
)

// These constants can be used to document the direction of a rotation, e.g.
// use Rotation(Clockwise * 0.5) to rotate clockwise by 0.5 (2 * math.Pi being a full rotation).
const (
	Clockwise        = 1.0
	CounterClockwise = -1.0
)

// FullAngle is a full rotation, i.e. 360°. Useful to document something like  quarter
// rotation which can be written as FullAngle/4.
const FullAngle = 2 * math.Pi

// Rotation creates a transformation which rotates vectors by the given angle.
// The rotation is clockwise if angle > 0 and clockwise if angle < 0.
func Rotation(angle float64) Transformation {
	return transformationMatrix{
		math.Cos(angle), -math.Sin(angle), 0.0,
		math.Sin(angle), math.Cos(angle), 0.0,
	}
}
