package ld45

type Object struct {
	// X is the logical X coordinate of the object in the current view.
	X int

	// Y is the logical Y coordinate of the object in the current view.
	Y int

	// Key is the object's identifier.
	Key string

	// Lifetime in ms.
	Lifetime int

	// GroundBound determines wether this Object should be rendered horizintally
	// centered and vertically bottom-aligned.
	GroundBound bool
}

type Objects []Object

func (objs Objects) Len() int {
	return len(objs)
}

func (objs Objects) Less(i, j int) bool {
	objI, objJ := objs[i], objs[j]
	if objI.GroundBound && !objJ.GroundBound {
		return true
	}
	if !objI.GroundBound && objJ.GroundBound {
		return false
	}
	return objs[i].Y < objs[j].Y
}

func (objs Objects) Swap(i, j int) {
	objs[i], objs[j] = objs[j], objs[i]
}
