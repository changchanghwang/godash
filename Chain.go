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
	return c.data[0]
}
