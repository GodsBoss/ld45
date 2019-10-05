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
}
