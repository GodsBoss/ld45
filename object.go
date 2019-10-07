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

var noObjects = make([]Object, 0)

type singleObject struct {
	sox   float64
	soy   float64
	cache []Object
}

func createSingleObject(x, y float64, groundBound bool) singleObject {
	return singleObject{
		sox: x,
		soy: y,
		cache: []Object{
			{
				GroundBound: groundBound,
			},
		},
	}
}

func (so *singleObject) ToObjects(cam camera) []Object {
	x, y := calculateScreenPosition(cam, so.sox, so.soy)
	so.cache[0].X = x
	so.cache[0].Y = y
	return so.cache
}

func (so *singleObject) setLifetime(lifetime int) {
	so.cache[0].Lifetime = lifetime
}

func (so *singleObject) setKey(key string) {
	so.cache[0].Key = key
}
