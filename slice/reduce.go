package slice

func Reduce[T any, Res any](src []T, init Res, reducer func(Res, T) Res) Res {
	for _, val := range src {
		init = reducer(init, val)
	}
	return init
}