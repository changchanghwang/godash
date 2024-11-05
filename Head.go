package godash

func Head[T comparable](target []T) T {
	var zero T
	if len(target) == 0 {
		return zero
	}

	return target[0]
}
