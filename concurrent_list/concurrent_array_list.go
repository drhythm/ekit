package concurrent_list

import (
	"github.com/drhythm/ekit/list"
	"sync"
)

// ConcurrentArrayList is a thread-safe implementation of an array list.
// It uses a Read-Write Mutex to ensure concurrency safety.
type ConcurrentArrayList[T any] struct {
	// The undering plain ArrayList.
	list	*list.ArrayList[T]
	// lock is a Read-Write Mutex.
	lock	sync.RWMutex
}

// NewConcurrentArrayList creates a new thread-safe ArrayList with the specified capacity.
func NewConcurrentArrayList[T any](cap int) *ConcurrentArrayList[T] {
	return &ConcurrentArrayList[T]{
		list:	list.NewArrayList[T](cap),

		// sync.RWMutex is initialized to its zero value automatically, so no need to set it manually.
	}
}

// Append adds an element to the end of list safely.
// This is a write operation, so it requires a Write Lock.
func (c *ConcurrentArrayList[T]) Append(val T) {
	// Write Lock
	c.lock.Lock()
	// Unlock
	defer c.lock.Unlock()
	// Delegate: Call the underlying list's Append method.
	c.list.Append(val)
}

// Delete removes the element at the specified index safety.
// This modifies the list, so it requires a Write Lock.
func (c *ConcurrentArrayList[T]) Delete(index int) (T, error) {
	// Write Lock
	c.lock.Lock()
	// Unlock
	defer c.lock.Unlock()
	// Delegate: Call the underlying list's Delete method.
	return c.list.Delete(index)
}

// Get retrieves the element at the specified index safety.
// This is a read operation, so it requires a Read Lock.
func (c *ConcurrentArrayList[T]) Get(index int) (T, error) {
	// Write Lock
	c.lock.RLock()
	// Unlock
	defer c.lock.Unlock()
	// Delegate: Call the underlying list's Get method.
	return c.list.Get(index)
}

// Len returns the number of elements in the list safety.
// Len is a read operation, so it requires a Read Lock.
func (c *ConcurrentArrayList[T]) Len() int {
	// Write Lock
	c.lock.RLock()
	// Unlock
	defer c.lock.Unlock()
	// Delegate: Call the underlying list's Len method.
	return c.list.Len()
}

