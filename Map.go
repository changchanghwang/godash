package godash

func Map[T comparable, R any](target []T, fn func(T) R) []R {
	copied := make([]R, len(target))
	for i, v := range target {
		copied[i] = fn(v)
	}
	return copied
}
