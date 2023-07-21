// Package safepool provides a wrapper for sync.Pool which ensures type and state safety.
//
// Because the type constraint on the pool, you can always rely to retrieve the expected
// type specified to the pool. Also, all objects put into the pool need to implement
// ResetState which sets the state of the objects to a clean "zero" state on putting it
// back into the pool.
package safepool

import (
	"sync"
)

// ResetState describes an object which state can be reset to its "zero" state to be put back
// into the object pool ready for reuse.
type ResetState interface {
	ResetState()
}

// SafePool wraps sync.Pool in a type and state safe manner. You can always expect to retrieve
// the object type set to the pool on retrival due to the type constraint on the pool.
//
// All objects need to implement ResetState, which is called when an object instance is put back
// into the pool to ensure a clean "zero" state on retrival.
type SafePool[T ResetState] struct {
	p *sync.Pool
}

// New creates a new SafePool with the given new function to create new objects on demand.
func New[T ResetState](new func() T) SafePool[T] {
	return SafePool[T]{
		p: &sync.Pool{
			New: func() any { return new() },
		},
	}
}

// Get retrieves an arbitrary object from the internal pool and removes it from the pool.
//
// For more information, read the documentation of sync.Pool#Get.
func (t SafePool[T]) Get() T {
	return t.p.Get().(T)
}

// Put first calls ResetState on the object and then puts it back into the pool.
func (t SafePool[T]) Put(v T) {
	v.ResetState()
	t.p.Put(v)
}
