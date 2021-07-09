//+build go1.9

package concurrent

import "sync"

<<<<<<< HEAD
=======
// Map is a wrapper for sync.Map introduced in go1.9
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
type Map struct {
	sync.Map
}

<<<<<<< HEAD
func NewMap() *Map {
	return &Map{}
}
=======
// NewMap creates a thread safe Map
func NewMap() *Map {
	return &Map{}
}
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
