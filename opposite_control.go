package ld45

// oppositeControl represents a pair of contradictionary controls which can still
// both be active, e.g. pressing left and right at the same time. Both can be
// enabled or disabled separately, but are merged into a single property.
type oppositeControl [2]bool

func (ctrl *oppositeControl) enableFirst() {
	ctrl[0] = true
}

func (ctrl *oppositeControl) disableFirst() {
	ctrl[0] = false
}

func (ctrl *oppositeControl) enableSecond() {
	ctrl[1] = true
}

func (ctrl *oppositeControl) disableSecond() {
	ctrl[1] = false
}

// isFirst returns wether only the first is enabled.
func (ctrl oppositeControl) isFirst() bool {
	return ctrl[0] && !ctrl[1]
}

// isSecond returns wether only the second is enabled.
func (ctrl oppositeControl) isSecond() bool {
	return ctrl[1] && !ctrl[0]
}

// isSome returns wether only the first or only the second is enabled.
func (ctrl oppositeControl) isSome() bool {
	return ctrl[0] != ctrl[1]
}

// isNone returns wether both are enabled or both are disabled.
func (ctrl oppositeControl) isNone() bool {
	return ctrl[0] == ctrl[1]
}

// asInt returns:
// 1 if only the first control is enabled.
// -1 if only the second control is enabled.
// 0 if they are both enabled or both disabled.
func (ctrl oppositeControl) asInt() int {
	return boolToInt[ctrl[0]] - boolToInt[ctrl[1]]
}

// asFloat64 works like asInt, but returns a float64 instead of an int.
func (ctrl oppositeControl) asFloat64() float64 {
	return boolToFloat64[ctrl[0]] - boolToFloat64[ctrl[1]]
}

var boolToInt = map[bool]int{
	false: 0,
	true:  1,
}

var boolToFloat64 = map[bool]float64{
	false: 0.0,
	true:  1.0,
}
