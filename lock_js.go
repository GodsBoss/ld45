// +build js

package ld45

// nopLock implements the same API as *sync.Mutex, but does nothing.
type nopLock struct{}

func (nop nopLock) Lock() {}

func (nop nopLock) Unlock() {}

func createLock() lock {
	return nopLock{}
}
