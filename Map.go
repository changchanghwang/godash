package godash

// Map returns a new slice with the results of applying the function to each element.
//
// @example
//
// var arr = []int{1, 2, 3, 4, 5}
//
// var result = Map(arr, func(v int) int {
// return v * 2
// }) // result is []int{2, 4, 6, 8, 10}
func Map[T comparable, R any](target []T, fn func(T) R) []R {
	copied := make([]R, len(target))
	for i, v := range target {
		copied[i] = fn(v)
	}
	return copied
}
