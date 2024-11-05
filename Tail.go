package godash

// Tail returns a new slice with the elements except the first element.
//
// @example
//
// var arr = []int{1, 2, 3}
//
// var result = Tail(arr) // result is []int{2, 3}
func Tail[T comparable](target []T) []T {
	if len(target) == 0 {
		return target
	}

	return target[1:]
}
