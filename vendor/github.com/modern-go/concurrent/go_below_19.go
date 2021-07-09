//+build !go1.9

package concurrent

import "sync"

<<<<<<< HEAD
=======
// Map implements a thread safe map for go version below 1.9 using mutex
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
type Map struct {
	lock sync.RWMutex
	data map[interface{}]interface{}
}

<<<<<<< HEAD
=======
// NewMap creates a thread safe map
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
func NewMap() *Map {
	return &Map{
		data: make(map[interface{}]interface{}, 32),
	}
}

<<<<<<< HEAD
=======
// Load is same as sync.Map Load
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
func (m *Map) Load(key interface{}) (elem interface{}, found bool) {
	m.lock.RLock()
	elem, found = m.data[key]
	m.lock.RUnlock()
	return
}

<<<<<<< HEAD
=======
// Load is same as sync.Map Store
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
func (m *Map) Store(key interface{}, elem interface{}) {
	m.lock.Lock()
	m.data[key] = elem
	m.lock.Unlock()
}
<<<<<<< HEAD

=======
>>>>>>> 9362ae084505e4d2b7e6c8fa897cf6dfdb8d64f7
