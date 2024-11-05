package godash

func Tail[T comparable](target []T) T {
	var zero T

	if len(target) == 0 {
		return zero
	}

	return target[len(target)-1]
}
