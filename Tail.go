package godash

func Tail[T comparable](target []T) []T {
	if len(target) == 0 {
		return target
	}

	return target[1:]
}
