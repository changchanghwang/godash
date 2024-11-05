package godash

type GodashChain[T comparable] struct {
	data []T
}

// Chain returns a new GodashChain instance. It is useful to chain multiple operations.
//
// @example
//
// var arr = []int{1, 2, 3}
//
// var result = godash.Chain(arr).Map(func(v int) int {
// return v * 2
// }).Filter(func(v int) bool {
// return v%2 == 0
// }).Value() // result is []int{2,4,6}
func Chain[T comparable](data []T) *GodashChain[T] {
	return &GodashChain[T]{
		data: DeepCopySlice(data),
	}
}

// Value returns the slice.
//
// @example
//
// var arr = []int{1, 2, 3}
//
// var result = godash.Chain(arr).Value() // result is []int{1, 2, 3}
func (c *GodashChain[T]) Value() []T {
	return c.data
}

// Head returns the first element of the slice.
//
// @example
//
// var arr = []int{1, 2, 3}
//
// var result, ok = godash.Chain(arr).Head() // result is 1, ok is true
func (c *GodashChain[T]) Head() (T, bool) {
	return Head(c.data)
}

// Tail returns a new slice with the elements except the first element.
//
// @example
//
// var arr = []int{1, 2, 3}
//
// var result = godash.Chain(arr).Tail().Value() // result is []int{2, 3}
func (c *GodashChain[T]) Tail() *GodashChain[T] {
	return Chain(Tail(c.data))
}

// Map returns a new slice with the results of applying the function to each element.
//
// @example
//
// var arr = []int{1, 2, 3, 4, 5}
//
// var result = godash.Chain(arr).Map(func(v int) int {
// return v * 2
// }).Value() // result is []int{2, 4, 6, 8, 10}
//
// @caution The Map of chain only works with the same type of the original slice.
func (c *GodashChain[T]) Map(fn func(T) T) *GodashChain[T] {
	return Chain(Map(c.data, fn))
}

// MapToInt returns a new slice with the results of applying the function to each element.
//
// @example
//
// var arr = []string{"1", "2", "3", "4", "5"}
//
// var result = godash.Chain(arr).MapToInt(func(v string) int {
// converted, err := strconv.Atoi(v)
//
//	if err != nil {
//	  return 0
//	}
//
// return converted * 2
// }).Value() // result is []int{2, 4, 6, 8, 10}
func (c *GodashChain[T]) MapToInt(fn func(T) int) *GodashChain[int] {
	return Chain(Map(c.data, fn))
}

// MapToBool returns a new slice with the results of applying the function to each element.
//
// @example
//
// var arr = []int{1, 2, 3, 4, 5}
//
// var result = godash.Chain(arr).MapToBool(func(v int) bool {
// return v%2 == 0
// }).Value() // result is []bool{false, true, false, true, false}
func (c *GodashChain[T]) MapToBool(fn func(T) bool) *GodashChain[bool] {
	return Chain(Map(c.data, fn))
}

// MapToString returns a new slice with the results of applying the function to each element.
//
// @example
//
// var arr = []int{1, 2, 3, 4, 5}
//
// var result = godash.Chain(arr).MapToString(func(v int) string {
// return strconv.Itoa(v)
// }).Value() // result is []string{"1", "2", "3", "4", "5"}
func (c *GodashChain[T]) MapToString(fn func(T) string) *GodashChain[string] {
	return Chain(Map(c.data, fn))
}

// Filter returns a new slice with the elements that pass the condition.
//
// @example
//
// var arr = []int{1, 2, 3, 4, 5}
//
// var result = godash.Chain(arr).Filter(func(v int) bool {
// return v%2 == 0
// }).Value() // result is []int{2, 4}
func (c *GodashChain[T]) Filter(fn func(T) bool) *GodashChain[T] {
	return Chain(Filter(c.data, fn))
}
