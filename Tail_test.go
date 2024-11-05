package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestTail(t *testing.T) {
	t.Run("Without Chain struct", func(t *testing.T) {
		t.Run("should return the elements except first of the slice", func(t *testing.T) {
			result := godash.Tail([]int{1, 2, 3})
			assert.Equal(t, []int{2, 3}, result)
		})

		t.Run("should return empty slice when input is empty", func(t *testing.T) {
			result := godash.Tail([]int{})
			assert.Equal(t, []int{}, result)
		})
	})

	t.Run("With Chain struct", func(t *testing.T) {
		t.Run("should return the elements except first of the chaining slice", func(t *testing.T) {
			result := godash.Chain([]int{}).Tail().Value()
			assert.Equal(t, []int{}, result)
		})
	})
}
