package godash

func Tail[T comparable](target []T) (T, bool) {
	var zero T

	if len(target) == 0 {
		return zero, false
	}

	return target[len(target)-1], true
}
