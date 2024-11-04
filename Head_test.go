package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	t.Run("It should return the first element of the slice", func(t *testing.T) {
		target := []int{1, 2, 3}
		actual := godash.Head(target)
		assert.Equal(t, 1, actual)
	})

	t.Run("It should return the first element of the chaining slice", func(t *testing.T) {
		target := []string{"A", "B", "C"}
		actual := godash.Chain(target).Head()

		assert.Equal(t, "A", actual)
	})
}
