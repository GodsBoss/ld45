package coords_test

import (
	"github.com/GodsBoss/ld45/pkg/coords"

	"testing"
)

func TestRotation(t *testing.T) {
	runTestCases(
		t,
		map[string]testCase{
			"no_rotation": rotationTestCase{
				rotation:  0,
				input:     newVector(7.0, 3.0),
				expected:  newVector(7.0, 3.0),
				tolerance: sameTolerance(0.0001),
			},
			"clockwise_1/8": rotationTestCase{
				rotation:  coords.Clockwise * coords.FullAngle * 0.125,
				input:     newVector(-3.0, 0),
				expected:  newVector(-2.1213203435596424, -2.1213203435596424),
				tolerance: sameTolerance(0.0001),
			},
			"clockwise_7/8": rotationTestCase{
				rotation:  coords.Clockwise * coords.FullAngle * 0.875,
				input:     newVector(2.4748737341529163, -2.4748737341529163),
				expected:  newVector(0, -3.5),
				tolerance: sameTolerance(0.0001),
			},
			"counter-clockwise_1/4": rotationTestCase{
				rotation:  coords.CounterClockwise * coords.FullAngle * 0.25,
				input:     newVector(-2.5, 2.5),
				expected:  newVector(2.5, 2.5),
				tolerance: sameTolerance(0.0001),
			},
			"counter-clockwise_3/4": rotationTestCase{
				rotation:  coords.CounterClockwise * coords.FullAngle * 0.75,
				input:     newVector(3.75, -3.75),
				expected:  newVector(3.75, 3.75),
				tolerance: sameTolerance(0.0001),
			},
		},
	)
}

type rotationTestCase struct {
	rotation  float64
	input     testVector
	expected  testVector
	tolerance testVector
}

func (testCase rotationTestCase) run(name string, t *testing.T) {
	t.Run(
		name,
		func(t *testing.T) {
			input := coords.VectorFromCartesian(testCase.input.x, testCase.input.y)
			rotation := coords.Rotation(testCase.rotation)
			actual := rotation.Transform(input)
			assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
		},
	)
}
