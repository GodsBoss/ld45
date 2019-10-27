// +build !js

package ld45

import (
	"sync"
)

func createLock() lock {
	return new(sync.Mutex)
}
