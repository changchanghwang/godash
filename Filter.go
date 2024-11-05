package godash

func Filter[T comparable](target []T, fn func(T) bool) []T {
	copied := make([]T, 0)

	for _, v := range target {
		if fn(v) {
			copied = append(copied, v)
		}
	}

	return copied
}
