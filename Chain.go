package godash

type GodashChain[T comparable] struct {
	data []T
}

func Chain[T comparable](data []T) *GodashChain[T] {
	return &GodashChain[T]{
		data: DeepCopySlice(data),
	}
}

func (c *GodashChain[T]) Value() []T {
	return c.data
}

func (c *GodashChain[T]) Head() T {
	return Head(c.data)
}
func (c *GodashChain[T]) Tail() T {
	return Tail(c.data)
}

func (c *GodashChain[T]) Map(fn func(T) T) *GodashChain[T] {
	return Chain(Map(c.data, fn))
}

func (c *GodashChain[T]) MapToInt(fn func(T) int) *GodashChain[int] {
	return Chain(Map(c.data, fn))
}

func (c *GodashChain[T]) MapToBool(fn func(T) bool) *GodashChain[bool] {
	return Chain(Map(c.data, fn))
}

func (c *GodashChain[T]) MapToString(fn func(T) string) *GodashChain[string] {
	return Chain(Map(c.data, fn))
}
