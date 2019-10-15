package coords_test

import (
	"github.com/GodsBoss/ld45/pkg/coords"

	"testing"
)

func TestTranslation(t *testing.T) {
	runTestCases(
		t,
		map[string]testCase{
			"left_up": translationTestCase{
				translation: newVector(coords.Left*2.0, coords.Up*4.0),
				input:       newVector(8.5, 3.4),
				expected:    newVector(6.5, -0.6),
				tolerance:   sameTolerance(0.0001),
			},
			"left_down": translationTestCase{
				translation: newVector(coords.Left*3.5, coords.Down*8.0),
				input:       newVector(2.5, 1.0),
				expected:    newVector(-1.0, 9.0),
				tolerance:   sameTolerance(0.0001),
			},
			"right_up": translationTestCase{
				translation: newVector(coords.Right*5.0, coords.Up*3.5),
				input:       newVector(-2.5, 7.5),
				expected:    newVector(2.5, 4.0),
				tolerance:   sameTolerance(0.0001),
			},
			"right_down": translationTestCase{
				translation: newVector(coords.Right*1.5, coords.Down*2.5),
				input:       newVector(3.0, -1.0),
				expected:    newVector(4.5, 1.5),
				tolerance:   sameTolerance(0.0001),
			},
		},
	)
}

type translationTestCase struct {
	translation testVector
	input       testVector
	expected    testVector
	tolerance   testVector
}

func (testCase translationTestCase) run(name string, t *testing.T) {
	t.Run(
		name+" (cartesian)",
		func(t *testing.T) {
			input := coords.VectorFromCartesian(testCase.input.x, testCase.input.y)
			translation := coords.Translation(testCase.translation.x, testCase.translation.y)
			actual := translation.Transform(input)
			assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
		},
	)
	t.Run(
		name+" (vector)",
		func(t *testing.T) {
			input := coords.VectorFromCartesian(testCase.input.x, testCase.input.y)
			translation := coords.TranslationByVector(coords.VectorFromCartesian(testCase.translation.x, testCase.translation.y))
			actual := translation.Transform(input)
			assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
		},
	)
}
