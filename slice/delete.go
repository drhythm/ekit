package slice

import (
	"errors"
)

func Delete[T any](src []T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index >= length {
		return nil, errors.New("index out of bounds")
	}
	copy(src[index:], src[index + 1:])
	return src[:length - 1], nil
}