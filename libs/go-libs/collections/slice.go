package collections

func Filter[T any](slice []T, filter func(T) bool) []T {
	ret := make([]T, 0)
	for _, item := range slice {
		if filter(item) {
			ret = append(ret, item)
		}
	}
	return ret
}
