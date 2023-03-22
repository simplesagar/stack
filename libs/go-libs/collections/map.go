package collections

func MapKeys[K comparable, V any](v map[K]V) []K {
	ret := make([]K, 0)
	for k := range v {
		ret = append(ret, k)
	}
	return ret
}
