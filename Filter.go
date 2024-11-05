package godash

// Filter returns a new slice with the elements that pass the condition.
//
// @example
//
// var arr = []int{1, 2, 3, 4, 5}
//
// var result = Filter(arr, func(v int) bool {
// return v%2 == 0
// }) // result is []int{2, 4}
func Filter[T comparable](target []T, fn func(T) bool) []T {
	copied := make([]T, 0)

	for _, v := range target {
		if fn(v) {
			copied = append(copied, v)
		}
	}

	return copied
}
