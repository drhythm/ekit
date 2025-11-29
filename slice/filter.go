package slice

func Filter[T any](src []T, predicate func(T) bool) []T {
	res := make([]T, 0, len(src))
	for _, val := range src {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}