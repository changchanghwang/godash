package godash_test

import (
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestDeppCopySlice(t *testing.T) {
	t.Run("It should return a new slice with the same elements", func(t *testing.T) {
		target := []int{1, 2, 3}
		actual := godash.DeepCopySlice(target)

		assert.NotSame(t, target, actual)
		assert.Equal(t, target, actual)
	})
}
