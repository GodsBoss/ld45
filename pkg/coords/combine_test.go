package coords_test

import (
	"fmt"

	"github.com/GodsBoss/ld45/pkg/coords"

	"testing"
)

func TestCombine(t *testing.T) {
	runTestCases(
		t,
		map[string]testCase{
			"identity": combineTestCase{
				transformations: make([]coords.Transformation, 0),
				input:           newVector(2.5, 3.5),
				expected:        newVector(2.5, 3.5),
				tolerance:       sameTolerance(0.0001),
			},
			"one_fixed": combineTestCase{
				transformations: []coords.Transformation{fixedTransformation(newVector(1.8, -2.4))},
				input:           newVector(-3.4, 5.8),
				expected:        newVector(1.8, -2.4),
				tolerance:       sameTolerance(0.0001),
			},
			"three_combined": combineTestCase{
				transformations: []coords.Transformation{
					coords.Translation(2.5, -3.0),
					coords.Scale(2.0, 1.5),
					coords.Rotation(coords.FullAngle * 0.25 * coords.Clockwise),
				},
				input:     newVector(3.0, 2.5),
				expected:  newVector(-2.5, 1.5),
				tolerance: sameTolerance(0.0001),
			},
			"custom_transformation_fixed_in_between": combineTestCase{
				transformations: []coords.Transformation{
					coords.Translation(-3.0, 4.5),
					coords.Scale(2.0, 3.0),
					fixedTransformation(newVector(1.8, -2.4)),
					coords.Scale(1.2, 2.3),
					coords.Translation(-3.0, 4.1),
					coords.Rotation(-1.2),
				},
				input:     newVector(-3.4, 5.8), // Irrelevant, because fixedTransformation ignores input
				expected:  newVector(0.6, -2.7),
				tolerance: sameTolerance(0.0001),
			},
			"custom_transformation_id_first": combineTestCase{
				transformations: []coords.Transformation{
					coords.Translation(-0.4, 0.8),
					testIdentity{},
				},
				input:     newVector(-1.4, 5.2),
				expected:  newVector(-1.8, 6.0),
				tolerance: sameTolerance(0.0001),
			},
			"custom_transformation_id_last": combineTestCase{
				transformations: []coords.Transformation{
					testIdentity{},
					coords.Translation(-0.4, 0.8),
				},
				input:     newVector(-1.4, 5.2),
				expected:  newVector(-1.8, 6.0),
				tolerance: sameTolerance(0.0001),
			},
		},
	)
}

type combineTestCase struct {
	transformations []coords.Transformation
	input           testVector
	expected        testVector
	tolerance       testVector
}

func (testCase combineTestCase) run(name string, t *testing.T) {
	combinedTransformation := coords.Combine(testCase.transformations...)
	fmt.Printf("combined_transformation = %#v\n", combinedTransformation)
	actual := combinedTransformation.Transform(coords.VectorFromCartesian(testCase.input.x, testCase.input.y))
	assertVectorWithin(t, name, testCase.expected, testCase.tolerance, actual)
}

type fixedTransformation testVector

func (ft fixedTransformation) Transform(_ coords.Vector) coords.Vector {
	return coords.VectorFromCartesian(ft.x, ft.y)
}

type testIdentity struct{}

func (ti testIdentity) Transform(v coords.Vector) coords.Vector {
	return v
}
