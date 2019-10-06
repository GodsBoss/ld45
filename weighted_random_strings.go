package ld45

import "math/rand"

type weightedRandomStrings map[string]int

func (wrs weightedRandomStrings) Random() string {
	full := 0
	for key := range wrs {
		full += wrs[key]
	}
	r := rand.Intn(full)
	until := 0
	for key := range wrs {
		until += wrs[key]
		if r < until {
			return key
		}
	}
	return ""
}
