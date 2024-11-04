package godash

func Map[T comparable](target []T, fn func(T) T) []T {
	copied := DeepCopySlice(target)
	for i, v := range target {
		copied[i] = fn(v)
	}
	return copied
}
