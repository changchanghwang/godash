package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	t.Run("Without Chain", func(t *testing.T) {
		t.Run("It should return a new slice with filtered slice", func(t *testing.T) {
			arr := []int{1, 2, 3, 4, 5}
			result := godash.Filter(arr, func(v int) bool {
				return v%2 == 0
			})

			assert.Equal(t, []int{2, 4}, result)
			assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
		})
	})

	t.Run("With Chain", func(t *testing.T) {
		t.Run("It should return a new slice with filtered slice", func(t *testing.T) {
			arr := []int{1, 2, 3, 4, 5}
			result := godash.Chain(arr).Filter(func(v int) bool {
				return v%2 == 0
			}).Value()

			assert.Equal(t, []int{2, 4}, result)
			assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
		})
	})
}
