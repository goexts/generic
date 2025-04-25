package cmp

import (
	"testing"
)

func TestIsZero(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  bool
	}{
		// Numeric types
		{"int zero", 0, true},
		{"int non-zero", 42, false},
		{"float64 zero", 0.0, true},
		{"float64 non-zero", 3.14, false},
		{"complex128 zero", complex(0, 0), true},
		{"complex128 non-zero", complex(1, 1), false},

		// String type
		{"string empty", "", true},
		{"string non-empty", "hello", false},

		// Boolean type
		{"bool false", false, true},
		{"bool true", true, false},

		// Pointer types
		{"nil pointer", (*int)(nil), true},
		{"non-nil pointer", new(int), false},

		// Struct types
		{"empty struct", struct{}{}, true},
		{
			name: "non-zero struct",
			input: struct {
				Name string
				Age  int
			}{Name: "Alice", Age: 30},
			want: false,
		},
		{
			name: "partially zero struct",
			input: struct {
				Name string
				Age  int
			}{Name: "", Age: 30},
			want: false,
		},

		// Array types
		{"zero array", [3]int{}, true},
		{"non-zero array", [3]int{0, 0, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case int:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case float64:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case complex128:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case string:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case bool:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case *int:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case struct{}:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			case [3]int:
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			default:
				// For custom struct types
				if got := IsZero(v); got != tt.want {
					t.Errorf("IsZero() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestZeroOr(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		defaults interface{}
		expected interface{}
	}{
		// int tests
		{
			name:     "int zero value returns default",
			input:    0,
			defaults: 42,
			expected: 42,
		},
		{
			name:     "int non-zero value returns input",
			input:    10,
			defaults: 42,
			expected: 10,
		},

		// string tests
		{
			name:     "empty string returns default",
			input:    "",
			defaults: "default",
			expected: "default",
		},
		{
			name:     "non-empty string returns input",
			input:    "hello",
			defaults: "default",
			expected: "hello",
		},

		// float tests
		{
			name:     "float zero value returns default",
			input:    0.0,
			defaults: 3.14,
			expected: 3.14,
		},
		{
			name:     "float non-zero value returns input",
			input:    1.23,
			defaults: 3.14,
			expected: 1.23,
		},

		// bool tests
		{
			name:     "bool false returns default",
			input:    false,
			defaults: true,
			expected: true,
		},
		{
			name:     "bool true returns input",
			input:    true,
			defaults: false,
			expected: true,
		},

		// struct tests
		{
			name:     "nil struct returns default",
			input:    (*struct{})(nil),
			defaults: &struct{}{},
			expected: &struct{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch v := tt.input.(type) {
			case int:
				result := ZeroOr(v, tt.defaults.(int))
				if result != tt.expected {
					t.Errorf("ZeroOr(%v, %v) = %v, want %v", v, tt.defaults, result, tt.expected)
				}
			case string:
				result := ZeroOr(v, tt.defaults.(string))
				if result != tt.expected {
					t.Errorf("ZeroOr(%v, %v) = %v, want %v", v, tt.defaults, result, tt.expected)
				}
			case float64:
				result := ZeroOr(v, tt.defaults.(float64))
				if result != tt.expected {
					t.Errorf("ZeroOr(%v, %v) = %v, want %v", v, tt.defaults, result, tt.expected)
				}
			case bool:
				result := ZeroOr(v, tt.defaults.(bool))
				if result != tt.expected {
					t.Errorf("ZeroOr(%v, %v) = %v, want %v", v, tt.defaults, result, tt.expected)
				}
			case *struct{}:
				result := ZeroOr(v, tt.defaults.(*struct{}))
				if result != tt.expected {
					t.Errorf("ZeroOr(%v, %v) = %v, want %v", v, tt.defaults, result, tt.expected)
				}
			}
		})
	}
}

func TestOr(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: 0, // zero value for int
		},
		{
			name:     "single value",
			input:    []int{42},
			expected: 42,
		},
		{
			name:     "multiple values, first non-zero",
			input:    []int{1, 2, 3},
			expected: 1,
		},
		{
			name:     "multiple values with zero, first non-zero",
			input:    []int{1, 0, 3},
			expected: 1,
		},
		{
			name:     "all zeros",
			input:    []int{0, 0, 0},
			expected: 0,
		},
		{
			name:     "first zero, second non-zero",
			input:    []int{0, 2, 3},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Or(tt.input...)
			if got != tt.expected {
				t.Errorf("Or(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}

func TestOrString(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "empty string first",
			input:    []string{"", "hello", "world"},
			expected: "hello",
		},
		{
			name:     "all empty strings",
			input:    []string{"", "", ""},
			expected: "",
		},
		{
			name:     "non-empty first",
			input:    []string{"first", "second"},
			expected: "first",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Or(tt.input...)
			if got != tt.expected {
				t.Errorf("Or(%v) = %v, want %v", tt.input, got, tt.expected)
			}
		})
	}
}
