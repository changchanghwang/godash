package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestTail(t *testing.T) {
	t.Run("Without Chain struct", func(t *testing.T) {
		t.Run("should return the last element of the slice", func(t *testing.T) {
			result, exist := godash.Tail([]int{1, 2, 3})
			assert.Equal(t, 3, result)
			assert.Equal(t, true, exist)
		})

		t.Run("should return nil of empty slice", func(t *testing.T) {
			result, exist := godash.Tail([]int{})
			assert.Equal(t, 0, result)
			assert.Equal(t, false, exist)
		})
	})

	t.Run("With Chain struct", func(t *testing.T) {
		t.Run("should return the last element of the chaining slice", func(t *testing.T) {
			result, exist := godash.Chain([]int{}).Tail()
			assert.Equal(t, 0, result)
			assert.Equal(t, false, exist)
		})
	})
}
