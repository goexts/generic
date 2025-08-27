package generic

import (
	"testing"
)

func TestHas(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		element  int
		expected bool
	}{
		{
			name:     "element present",
			slice:    []int{1, 2, 3, 4, 5},
			element:  3,
			expected: true,
		},
		{
			name:     "element not present",
			slice:    []int{1, 2, 3, 4, 5},
			element:  6,
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			element:  1,
			expected: false,
		},
		{
			name:     "first element",
			slice:    []int{10, 20, 30},
			element:  10,
			expected: true,
		},
		{
			name:     "last element",
			slice:    []int{10, 20, 30},
			element:  30,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Has(tt.slice, tt.element)
			if result != tt.expected {
				t.Errorf("Has(%v, %d) = %v, want %v", tt.slice, tt.element, result, tt.expected)
			}
		})
	}
}

func TestHas_StringType(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		expected bool
	}{
		{
			name:     "string present",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "banana",
			expected: true,
		},
		{
			name:     "string not present",
			slice:    []string{"apple", "banana", "cherry"},
			element:  "orange",
			expected: false,
		},
		{
			name:     "empty string",
			slice:    []string{"", "hello", "world"},
			element:  "",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Has(tt.slice, tt.element)
			if result != tt.expected {
				t.Errorf("Has(%v, %q) = %v, want %v", tt.slice, tt.element, result, tt.expected)
			}
		})
	}
}

func TestHas_BoolType(t *testing.T) {
	tests := []struct {
		name     string
		slice    []bool
		element  bool
		expected bool
	}{
		{
			name:     "true present",
			slice:    []bool{true, false, true},
			element:  true,
			expected: true,
		},
		{
			name:     "false present",
			slice:    []bool{true, false, true},
			element:  false,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []bool{true, true},
			element:  false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Has(tt.slice, tt.element)
			if result != tt.expected {
				t.Errorf("Has(%v, %v) = %v, want %v", tt.slice, tt.element, result, tt.expected)
			}
		})
	}
}
