package maps

import (
	"fmt"
	"strings"
	"testing"
)

func TestFirstKey(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		keys      []string
		wantKey   string
		wantFound bool
	}{
		{
			name:      "Key exists at first position",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{"apple", "orange"},
			wantKey:   "apple",
			wantFound: true,
		},
		{
			name:      "Key exists at later position",
			m:         map[string]int{"apple": 1, "banana": 2, "grape": 3},
			keys:      []string{"orange", "banana", "kiwi"},
			wantKey:   "banana",
			wantFound: true,
		},
		{
			name:      "No key exists",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{"orange", "kiwi"},
			wantKey:   "", // Zero value for string
			wantFound: false,
		},
		{
			name:      "Empty key list",
			m:         map[string]int{"apple": 1},
			keys:      []string{},
			wantKey:   "",
			wantFound: false,
		},
		{
			name:      "Empty map",
			m:         map[string]int{},
			keys:      []string{"apple", "banana"},
			wantKey:   "",
			wantFound: false,
		},
		{
			name:      "Single key in list, exists",
			m:         map[string]int{"apple": 1},
			keys:      []string{"apple"},
			wantKey:   "apple",
			wantFound: true,
		},
		{
			name:      "Single key in list, does not exist",
			m:         map[string]int{"apple": 1},
			keys:      []string{"banana"},
			wantKey:   "",
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotFound := FirstKey(tt.m, tt.keys...)
			if gotKey != tt.wantKey || gotFound != tt.wantFound {
				t.Errorf("FirstKey() got = (%v, %v), want (%v, %v)", gotKey, gotFound, tt.wantKey, tt.wantFound)
			}
		})
	}
}

func TestFirstKeyBy(t *testing.T) {
	// Example transform: case-insensitive string comparison
	toLowerCase := func(s string) string {
		return strings.ToLower(s)
	}

	tests := []struct {
		name      string
		m         map[string]int // Map keys are consistently string in these tests
		keys      interface{}    // Can be []string or []int
		transform interface{}    // Can be func(string) string or func(int) string
		wantKey   interface{}    // Can be string or int
		wantFound bool
	}{
		{
			name:      "Case-insensitive match at first position",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{"Apple", "Orange"},
			transform: toLowerCase,
			wantKey:   "Apple",
			wantFound: true,
		},
		{
			name:      "Case-insensitive match at later position",
			m:         map[string]int{"apple": 1, "banana": 2, "grape": 3},
			keys:      []string{"Orange", "BANANA", "Kiwi"},
			transform: toLowerCase,
			wantKey:   "BANANA",
			wantFound: true,
		},
		{
			name:      "No key exists after transform",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{"Orange", "Kiwi"},
			transform: toLowerCase,
			wantKey:   "",
			wantFound: false,
		},
		{
			name:      "Empty key list",
			m:         map[string]int{"apple": 1},
			keys:      []string{},
			transform: toLowerCase,
			wantKey:   "",
			wantFound: false,
		},
		{
			name:      "Empty map",
			m:         map[string]int{},
			keys:      []string{"Apple", "Banana"},
			transform: toLowerCase,
			wantKey:   "",
			wantFound: false,
		},
		{
			name:      "Single key in list, exists after transform",
			m:         map[string]int{"apple": 1},
			keys:      []string{"APPLE"},
			transform: toLowerCase,
			wantKey:   "APPLE",
			wantFound: true,
		},
		{
			name:      "Single key in list, does not exist after transform",
			m:         map[string]int{"apple": 1},
			keys:      []string{"Banana"},
			transform: toLowerCase,
			wantKey:   "",
			wantFound: false,
		},
		{
			name:      "Transform to different type (int to string)",
			m:         map[string]int{"10": 1, "20": 2},
			keys:      []int{5, 10, 15},
			transform: func(i int) string { return fmt.Sprintf("%d", i) },
			wantKey:   10,
			wantFound: true,
		},
		{
			name:      "Transform to different type (int to string), no match",
			m:         map[string]int{"10": 1, "20": 2},
			keys:      []int{5, 15, 25},
			transform: func(i int) string { return fmt.Sprintf("%d", i) },
			wantKey:   0, // Zero value for int
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotKey interface{}
			var gotFound bool

			// Determine the type of keys and transform function using type assertion
			switch k := tt.keys.(type) {
			case []string:
				trans, ok := tt.transform.(func(string) string)
				if !ok {
					t.Fatalf("Test case %q: expected transform to be func(string) string, got %T", tt.name, tt.transform)
				}
				// Call FirstKeyBy with string keys and string transform (transform is now the second argument)
				stringGotKey, stringGotFound := FirstKeyBy(tt.m, trans, k...)
				gotKey = stringGotKey
				gotFound = stringGotFound
			case []int:
				trans, ok := tt.transform.(func(int) string)
				if !ok {
					t.Fatalf("Test case %q: expected transform to be func(int) string, got %T", tt.name, tt.transform)
				}
				// Call FirstKeyBy with int keys and int transform (transform is now the second argument)
				intGotKey, intGotFound := FirstKeyBy(tt.m, trans, k...)
				gotKey = intGotKey
				gotFound = intGotFound
			default:
				t.Fatalf("Test case %q: unsupported keys type %T", tt.name, tt.keys)
			}

			if fmt.Sprintf("%v", gotKey) != fmt.Sprintf("%v", tt.wantKey) || gotFound != tt.wantFound {
				t.Errorf("FirstKeyBy() got = (%v, %v), want (%v, %v)", gotKey, gotFound, tt.wantKey, tt.wantFound)
			}
		})
	}
}
