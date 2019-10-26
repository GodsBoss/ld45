package coords_test

import (
	"github.com/GodsBoss/ld45/pkg/coords"

	"testing"
)

func TestScale(t *testing.T) {
	runTestCases(
		t,
		map[string]testCase{
			"scale_1": scaleTestCase{
				scaleH:    -2.0,
				scaleV:    1.5,
				input:     newVector(3, 4.0),
				expected:  newVector(-6.0, 6.0),
				tolerance: sameTolerance(0.0001),
			},
			"scale_2": scaleTestCase{
				scaleH:    3.5,
				scaleV:    -2.5,
				input:     newVector(-2.0, -3.0),
				expected:  newVector(-7.0, 7.5),
				tolerance: sameTolerance(0.0001),
			},
		},
	)
}

type scaleTestCase struct {
	scaleH    float64
	scaleV    float64
	input     testVector
	expected  testVector
	tolerance testVector
}

func (testCase scaleTestCase) run(name string, t *testing.T) {
	scale := coords.Scale(testCase.scaleH, testCase.scaleV)
	actual := scale.Transform(coords.VectorFromCartesian(testCase.input.x, testCase.input.y))
	assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
}
