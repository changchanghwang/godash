package godash

// Head returns the first element of the slice.
//
// @example
//
// var arr = []int{1, 2, 3}
//
// var result, ok = Head(arr) // result is 1, ok is true
func Head[T comparable](target []T) (T, bool) {
	var zero T
	if len(target) == 0 {
		return zero, false
	}

	return target[0], true
}
