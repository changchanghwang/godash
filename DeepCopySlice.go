package godash

// DeepCopySlice returns a new slice with the same elements.
//
// @example
//
// var target = []int{1, 2, 3}
//
// var actual = DeepCopySlice(target) // actual is []int{1, 2, 3}
func DeepCopySlice[T comparable](in []T) []T {
	out := make([]T, len(in))
	copy(out, in)
	return out
}
