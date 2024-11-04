package godash

func DeepCopySlice[T comparable](in []T) []T {
	out := make([]T, len(in))
	copy(out, in)
	return out
}
