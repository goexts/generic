package generic

import (
	"reflect"
	"sort"
	"testing"
)

func TestDifference(t *testing.T) {
	tests := []struct {
		name     string
		slice1   []int // The filter list
		slice2   []int // The data source
		expected []int
	}{
		{
			name:     "basic exclusion",
			slice1:   []int{1, 3, 5},
			slice2:   []int{1, 2, 3, 4, 5, 6},
			expected: []int{2, 4, 6},
		},
		{
			name:     "data source with duplicates",
			slice1:   []int{1, 5},
			slice2:   []int{1, 2, 2, 3, 4, 4, 5},
			expected: []int{2, 3, 4},
		},
		{
			name:     "no common elements",
			slice1:   []int{1, 2, 3},
			slice2:   []int{4, 5, 6},
			expected: []int{4, 5, 6},
		},
		{
			name:     "all data is filtered out",
			slice1:   []int{1, 2, 3, 4},
			slice2:   []int{1, 2, 2, 3},
			expected: []int{},
		},
		{
			name:     "empty filter list (slice1)",
			slice1:   []int{},
			slice2:   []int{1, 2, 3, 1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty data source (slice2)",
			slice1:   []int{1, 2, 3},
			slice2:   []int{},
			expected: []int{},
		},
		{
			name:     "both slices are identical",
			slice1:   []int{1, 2, 3},
			slice2:   []int{1, 2, 3},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Difference(tt.slice1, tt.slice2)
			// Sort both slices for stable comparison, as map iteration order is not guaranteed.
			sort.Ints(result)
			sort.Ints(tt.expected)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Difference() test %s failed: got %v, want %v", tt.name, result, tt.expected)
			}
		})
	}
}
