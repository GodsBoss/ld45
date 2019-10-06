package ld45

type sliceLen int

func (l sliceLen) Offset(start, offset int) int {
	if l == 0 {
		return -1
	}
	index := (start + offset) % int(l)
	if index < 0 {
		return index + int(l)
	}
	return index
}
