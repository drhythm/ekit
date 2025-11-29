package slice

func Map[Src any, Dst any](src []Src, mapper func(Src) Dst) []Dst {
    res := make([]Dst, len(src))
	for index, val := range src {
		res[index] = mapper(val)
	}
	return res
}