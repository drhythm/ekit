package list

import "errors"

type ArrayList[T any] struct {
	vals []T
}


func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{
		vals: make([]T, 0, cap),
	}
}

func NewArrayListOf[T any](src []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: src,
	}
}

func (a *ArrayList[T]) Append(val T) {
	a.vals = append(a.vals, val)
}

func (a *ArrayList[T]) Get(index int) (T, error) {
	if index < 0 || index >= len(a.vals) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return a.vals[index], nil
}

// Delete removes the element at the spacified index and return it.
func (a *ArrayList[T]) Delete(index int) (T, error) {
	// 1. Bounds check: Ensure the index is within the valid range.
	length := len(a.vals)
	if index < 0 || index >= length {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	// 2. Capture the value: Save the element to it later.
	res := a.vals[index]

	// 3. Shift elements: Move elements from index+1 one step to the left 
	copy(a.vals[index:], a.vals[index + 1:])

	// 4. Cleanup: Reset the last element to the zero value.
	// This prevents memory leaks, especially if T is a pointer type.
	var zero T
	a.vals[length - 1] = zero
	
	// 5. Update length: Reduce the slice length by 1.
	a.vals = a.vals[:length - 1]

	// 6. Shrink check
	capacity := cap(a.vals)
	if capacity > 64 && length - 1 < capacity / 4{
		tmp := make([]T, length - 1, capacity / 2)
		copy(tmp, a.vals)
		a.vals = tmp
	}
	return res, nil
}