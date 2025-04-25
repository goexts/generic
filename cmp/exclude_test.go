package cmp

import (
	"testing"
)

func TestExcludes(t *testing.T) {
	tests := []struct {
		name     string
		src      []int
		ts       []int
		expected []int
	}{
		{
			name:     "empty source",
			src:      []int{},
			ts:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "empty test values",
			src:      []int{1, 2, 3},
			ts:       []int{},
			expected: []int{},
		},
		{
			name:     "no exclusions",
			src:      []int{1, 2, 3},
			ts:       []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "some exclusions",
			src:      []int{1, 2, 3},
			ts:       []int{1, 4, 2, 5, 3, 6},
			expected: []int{4, 5, 6},
		},
		{
			name:     "all excluded",
			src:      []int{1, 2, 3},
			ts:       []int{4, 5, 6},
			expected: []int{4, 5, 6},
		},
		{
			name:     "with duplicates",
			src:      []int{1, 2, 3},
			ts:       []int{1, 4, 2, 4, 5, 3, 6, 6},
			expected: []int{4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Excludes(tt.src, tt.ts...)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("At index %d: expected %d, got %d", i, tt.expected[i], v)
				}
			}
		})
	}
}
