package coords

// Combine takes several transfomations and combines them into a single transformations.
// Transformations created by Identity(), Rotation(), Scale() and Translation() are
// merged into a single operation when they appear subsequently in the input.
//
// The result of transforming via the combined transformation is the same as if
// all transformations were applied to a vector, beginning with the right-most
// transformation and ending with the left-most one.
func Combine(transformations ...Transformation) Transformation {
	i := 1
	for i < len(transformations) {
		currentAsMatrix, currentIsMatrix := transformations[i].(transformationMatrix)
		// transformations[i] cannot be merged, so the next candidates are i+1 and i+2.
		if !currentIsMatrix {
			i += 2
			continue
		}
		lastAsMatrix, lastIsMatrix := transformations[i-1].(transformationMatrix)
		// transformations[i-1] cannot be merged, so the next candidates are i and i+1.
		if !lastIsMatrix {
			i += 1
			continue
		}
		// Merge transformations[i-1] and transformations[i] into a single transformation.
		nextTransformations := transformations[:i-1]
		nextTransformations = append(nextTransformations, multiplyTwoTransformationMatrices(lastAsMatrix, currentAsMatrix))
		nextTransformations = append(nextTransformations, transformations[i+1:]...)
		transformations = nextTransformations
	}
	if len(transformations) == 0 {
		return Identity()
	}
	if len(transformations) == 1 {
		return transformations[0]
	}
	return concatenation(transformations)
}

// Identity returns the identity transformation, i.e. the transformation which returns a vector as it is.
func Identity() Transformation {
	return transformationMatrix{
		1, 0, 0,
		0, 1, 0,
	}
}

type concatenation []Transformation

func (conc concatenation) Transform(v Vector) Vector {
	for i := len(conc) - 1; i >= 0; i-- {
		v = conc[i].Transform(v)
	}
	return v
}
