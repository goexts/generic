package set_test

import (
	"testing"

	"github.com/goexts/generic/set"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		assert.True(t, set.Contains([]int{1, 2, 3}, 2))
		assert.False(t, set.Contains([]int{1, 2, 3}, 4))
		assert.False(t, set.Contains([]int{}, 1))
	})

	t.Run("string slice", func(t *testing.T) {
		assert.True(t, set.Contains([]string{"a", "b", "c"}, "b"))
		assert.False(t, set.Contains([]string{"a", "b", "c"}, "d"))
	})
}

func TestExists(t *testing.T) {
	t.Run("int slice with predicate", func(t *testing.T) {
		// Check for an even number
		assert.True(t, set.Exists([]int{1, 3, 5, 6}, func(i int) bool { return i%2 == 0 }))
		// Check for a number greater than 10
		assert.False(t, set.Exists([]int{1, 3, 5, 6}, func(i int) bool { return i > 10 }))
	})

	t.Run("string slice with predicate", func(t *testing.T) {
		// Check for a string with length > 3
		assert.False(t, set.Exists([]string{"a", "bcd", "ef"}, func(s string) bool { return len(s) > 3 }))
		// Check for a string with length > 3 (positive case)
		assert.True(t, set.Exists([]string{"a", "abcd", "ef"}, func(s string) bool { return len(s) > 3 }))
		// Check for the specific string "hello"
		assert.False(t, set.Exists([]string{"a", "bcd", "ef"}, func(s string) bool { return s == "hello" }))
	})
}

func TestIntersection(t *testing.T) {
	t.Run("int slices", func(t *testing.T) {
		s1 := []int{1, 2, 3, 4, 5}
		s2 := []int{0, 1, 6, 3}
		expected := []int{1, 3}
		assert.ElementsMatch(t, expected, set.Intersection(s1, s2))

		// No intersection
		s3 := []int{10, 11, 12}
		assert.Empty(t, set.Intersection(s1, s3))

		// With duplicates
		s4 := []int{1, 1, 2, 2, 3, 3}
		s5 := []int{1, 2, 4}
		expected2 := []int{1, 2}
		assert.ElementsMatch(t, expected2, set.Intersection(s4, s5))
	})

	t.Run("string slices", func(t *testing.T) {
		s1 := []string{"apple", "banana", "cherry"}
		s2 := []string{"banana", "date", "apple"}
		expected := []string{"apple", "banana"}
		assert.ElementsMatch(t, expected, set.Intersection(s1, s2))

		// Case sensitive
		s3 := []string{"Apple", "Banana", "Cherry"}
		s4 := []string{"apple", "banana"}
		assert.Empty(t, set.Intersection(s3, s4))
	})
}

func TestDifference(t *testing.T) {
	t.Run("int slices", func(t *testing.T) {
		// Basic exclusion
		assert.ElementsMatch(t, []int{2, 4, 6}, set.Difference([]int{1, 3, 5}, []int{1, 2, 3, 4, 5, 6}))

		// Data source with duplicates
		assert.ElementsMatch(t, []int{2, 3, 4}, set.Difference([]int{1, 5}, []int{1, 2, 2, 3, 4, 4, 5}))

		// No common elements
		assert.ElementsMatch(t, []int{4, 5, 6}, set.Difference([]int{1, 2, 3}, []int{4, 5, 6}))

		// All data is filtered out
		assert.Empty(t, set.Difference([]int{1, 2, 3, 4}, []int{1, 2, 2, 3}))

		// Empty filter list
		assert.ElementsMatch(t, []int{1, 2, 3}, set.Difference([]int{}, []int{1, 2, 3, 1}))

		// Empty data source
		assert.Empty(t, set.Difference([]int{1, 2, 3}, []int{}))
	})
}

func TestUnion(t *testing.T) {
	t.Run("int slices", func(t *testing.T) {
		// Basic union
		s1 := []int{1, 2, 3}
		s2 := []int{3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		assert.ElementsMatch(t, expected, set.Union(s1, s2))

		// With duplicates in input slices
		s3 := []int{1, 2, 2, 3}
		s4 := []int{3, 4, 4, 5}
		expected = []int{1, 2, 3, 4, 5}
		assert.ElementsMatch(t, expected, set.Union(s3, s4))

		// First slice is empty
		s5 := []int{}
		s6 := []int{1, 2, 3}
		expected = []int{1, 2, 3}
		assert.ElementsMatch(t, expected, set.Union(s5, s6))

		// Second slice is empty
		s7 := []int{1, 2, 3}
		s8 := []int{}
		expected = []int{1, 2, 3}
		assert.ElementsMatch(t, expected, set.Union(s7, s8))

		// Both slices are empty
		s9 := []int{}
		s10 := []int{}
		assert.Empty(t, set.Union(s9, s10))
	})

	t.Run("string slices", func(t *testing.T) {
		// Basic string union
		s1 := []string{"apple", "banana"}
		s2 := []string{"banana", "cherry"}
		expected := []string{"apple", "banana", "cherry"}
		assert.ElementsMatch(t, expected, set.Union(s1, s2))

		// Case sensitivity
		s3 := []string{"Apple", "Banana"}
		s4 := []string{"apple", "banana"}
		expected = []string{"Apple", "Banana", "apple", "banana"}
		assert.ElementsMatch(t, expected, set.Union(s3, s4))

		// With empty strings
		s5 := []string{"", "a", "b"}
		s6 := []string{"a", "b", "c", ""}
		expected = []string{"", "a", "b", "c"}
		assert.ElementsMatch(t, expected, set.Union(s5, s6))
	})
}
