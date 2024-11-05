package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	t.Run("Without Chain struct", func(t *testing.T) {
		t.Run("It should return the first element of the slice", func(t *testing.T) {
			target := []int{1, 2, 3}
			result, exist := godash.Head(target)
			assert.Equal(t, 1, result)
			assert.Equal(t, true, exist)
		})

		t.Run("It should return the first element of the chaining slice", func(t *testing.T) {
			target := []string{"A", "B", "C"}
			result, exist := godash.Chain(target).Head()

			assert.Equal(t, "A", result)
			assert.Equal(t, true, exist)
		})
	})

	t.Run("With Chain struct", func(t *testing.T) {
		t.Run("It should return the first element of the chaining slice", func(t *testing.T) {
			target := []string{}
			result, exist := godash.Chain(target).Head()

			assert.Equal(t, "", result)
			assert.Equal(t, false, exist)
		})
	})

}
