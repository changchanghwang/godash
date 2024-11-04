package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("It should return a new slice with the function applied to each element", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := godash.Map(arr, func(v int) int {
			return v * 2
		})

		assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
	})

	t.Run("It should return a new slice with the function applied to each Chaining element", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		result := godash.Chain(arr).Map(func(v int) int {
			return v * 2
		}).Value()

		assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
		assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
	})
}
