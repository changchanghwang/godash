package godash_test

import (
	"strconv"
	"testing"

	"github.com/changchanghwang/godash"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("Without Chain struct", func(t *testing.T) {
		t.Run("It should return a new slice with the function applied to each element", func(t *testing.T) {
			arr := []int{1, 2, 3, 4, 5}
			result := godash.Map(arr, func(v int) int {
				return v * 2
			})

			assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
			assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
		})

		t.Run("It should return a new slice with the function applied to each element - to other type", func(t *testing.T) {
			arr := []int{1, 2, 3, 4, 5}
			result := godash.Map(arr, func(v int) bool {
				return v%2 == 0
			})

			assert.Equal(t, []bool{false, true, false, true, false}, result)
			assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
		})
	})

	t.Run("With Chain struct", func(t *testing.T) {
		t.Run("Map", func(t *testing.T) {
			t.Run("It should return a new slice with the function applied to each Chaining element", func(t *testing.T) {
				arr := []int{1, 2, 3, 4, 5}
				result := godash.Chain(arr).Map(func(v int) int {
					return v * 2
				}).Value()

				assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
				assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
			})
		})

		t.Run("MapToInt", func(t *testing.T) {
			t.Run("It should return a new slice with the function applied to each Chaining element", func(t *testing.T) {
				arr := []string{"1", "2", "3", "4", "5"}
				result := godash.Chain(arr).MapToInt(func(v string) int {
					converted, err := strconv.Atoi(v)
					if err != nil {
						return 0
					}

					return converted * 2
				}).Value()

				assert.Equal(t, []int{2, 4, 6, 8, 10}, result)
				assert.Equal(t, []string{"1", "2", "3", "4", "5"}, arr)
			})
		})

		t.Run("MapToBool", func(t *testing.T) {
			t.Run("It should return a new slice with the function applied to each Chaining element", func(t *testing.T) {
				arr := []int{1, 2, 3, 4, 5}
				result := godash.Chain(arr).MapToBool(func(v int) bool {
					return v%2 == 0
				}).Value()

				assert.Equal(t, []bool{false, true, false, true, false}, result)
				assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
			})
		})

		t.Run("MapToString", func(t *testing.T) {
			t.Run("It should return a new slice with the function applied to each Chaining element", func(t *testing.T) {
				arr := []int{1, 2, 3, 4, 5}
				result := godash.Chain(arr).MapToString(func(v int) string {
					return strconv.Itoa(v * 2)
				}).Value()

				assert.Equal(t, []string{"2", "4", "6", "8", "10"}, result)
				assert.Equal(t, []int{1, 2, 3, 4, 5}, arr)
			})
		})
	})
}
