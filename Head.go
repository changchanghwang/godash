package godash

func Head[T comparable](target []T) (T, bool) {
	var zero T
	if len(target) == 0 {
		return zero, false
	}

	return target[0], true
}
