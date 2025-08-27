package generic

import (
	"testing"
)

func TestIncludes(t *testing.T) {
	tests := []struct {
		name     string
		src      []int
		targets  []int
		expected []int
	}{
		{
			name:     "empty source and targets",
			src:      []int{},
			targets:  []int{},
			expected: []int{},
		},
		{
			name:     "empty source",
			src:      []int{},
			targets:  []int{1, 2, 3},
			expected: []int{},
		},
		{
			name:     "empty targets",
			src:      []int{1, 2, 3},
			targets:  []int{},
			expected: []int{},
		},
		{
			name:     "all targets included",
			src:      []int{1, 2, 3, 4, 5},
			targets:  []int{1, 3, 5},
			expected: []int{1, 3, 5},
		},
		{
			name:     "some targets included",
			src:      []int{1, 2, 3, 4, 5},
			targets:  []int{0, 1, 6, 3},
			expected: []int{1, 3},
		},
		{
			name:     "no targets included",
			src:      []int{1, 2, 3, 4, 5},
			targets:  []int{0, 6, 7, 8},
			expected: []int{},
		},
		{
			name:     "duplicate targets",
			src:      []int{1, 2, 3, 4, 5},
			targets:  []int{1, 1, 2, 2, 6},
			expected: []int{1, 1, 2, 2},
		},
		{
			name:     "duplicate source elements",
			src:      []int{1, 1, 2, 2, 3, 3},
			targets:  []int{1, 2, 4},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Includes(tt.src, tt.targets...)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("at index %d, expected %v, got %v", i, tt.expected[i], result[i])
				}
			}
		})
	}
}

func TestIncludesWithStrings(t *testing.T) {
	tests := []struct {
		name     string
		src      []string
		targets  []string
		expected []string
	}{
		{
			name:     "string values",
			src:      []string{"apple", "banana", "cherry"},
			targets:  []string{"banana", "date", "apple"},
			expected: []string{"banana", "apple"},
		},
		{
			name:     "case sensitive",
			src:      []string{"Apple", "Banana", "Cherry"},
			targets:  []string{"apple", "banana"},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Includes(tt.src, tt.targets...)
			if len(result) != len(tt.expected) {
				t.Errorf("expected length %d, got %d", len(tt.expected), len(result))
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("at index %d, expected %v, got %v", i, tt.expected[i], result[i])
				}
			}
		})
	}
}
