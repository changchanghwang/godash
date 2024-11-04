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

func (c *GodashChain[T]) Map(fn func(T) T) *GodashChain[T] {
	for i, v := range c.data {
		c.data[i] = fn(v)
	}
	return c
}
