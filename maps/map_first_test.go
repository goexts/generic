package maps_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/maps"
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
			gotKey, gotFound := maps.FirstKey(tt.m, tt.keys...)
			if gotKey != tt.wantKey || gotFound != tt.wantFound {
				t.Errorf("FirstKey() got = (%v, %v), want (%v, %v)", gotKey, gotFound, tt.wantKey, tt.wantFound)
			}
		})
	}
}

func TestFirstEntry(t *testing.T) {
	tests := []struct {
		name      string
		m         map[string]int
		keys      []string
		wantKey   string
		wantValue int
		wantFound bool
	}{
		{
			name:      "Key exists at first position",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{"apple", "orange"},
			wantKey:   "apple",
			wantValue: 1,
			wantFound: true,
		},
		{
			name:      "Key exists at second position",
			m:         map[string]int{"apple": 1, "banana": 2, "orange": 3},
			keys:      []string{"pear", "banana", "apple"},
			wantKey:   "banana",
			wantValue: 2,
			wantFound: true,
		},
		{
			name:      "No key exists",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{"pear", "orange"},
			wantKey:   "",
			wantValue: 0,
			wantFound: false,
		},
		{
			name:      "Empty keys slice",
			m:         map[string]int{"apple": 1, "banana": 2},
			keys:      []string{},
			wantKey:   "",
			wantValue: 0,
			wantFound: false,
		},
		{
			name:      "Empty map",
			m:         map[string]int{},
			keys:      []string{"apple", "banana"},
			wantKey:   "",
			wantValue: 0,
			wantFound: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, found := maps.FirstEntry(tt.m, tt.keys...)
			if found != tt.wantFound {
				t.Errorf("FirstEntry() found = %v, want %v", found, tt.wantFound)
				return
			}
			if found && (got.Key != tt.wantKey || got.Val != tt.wantValue) {
				t.Errorf("FirstEntry() = %+v, want key=%v value=%v", got, tt.wantKey, tt.wantValue)
			}
		})
	}
}

func TestFirstEntryBy(t *testing.T) {
	// Test case-insensitive search
	t.Run("case-insensitive search", func(t *testing.T) {
		m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

		tests := []struct {
			name      string
			keys      []string
			wantKey   string
			wantValue int
			wantFound bool
		}{
			{
				name:      "match first key",
				keys:      []string{"apple", "BANANA", "cherry"},
				wantKey:   "apple", // Returns the original key from the input list
				wantValue: 1,
				wantFound: true,
			},
			{
				name:      "match second key",
				keys:      []string{"pear", "banana", "cherry"},
				wantKey:   "banana", // Returns the original key from the input list
				wantValue: 2,
				wantFound: true,
			},
			{
				name:      "no match",
				keys:      []string{"PEAR", "MANGO"},
				wantKey:   "",
				wantValue: 0,
				wantFound: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got, found := maps.FirstEntryBy(m, strings.ToLower, tt.keys...)
				if found != tt.wantFound {
					t.Errorf("FirstEntryBy() found = %v, want %v", found, tt.wantFound)
					return
				}
				if found && (got.Key != tt.wantKey || got.Val != tt.wantValue) {
					t.Errorf("FirstEntryBy() = %+v, want key=%v value=%v", got, tt.wantKey, tt.wantValue)
				}
			})
		}
	})

	// Test with custom key type
	t.Run("custom key type", func(t *testing.T) {
		type productID struct {
			id   int
			name string
		}

		m := map[int]string{
			1001: "Laptop",
			2002: "Phone",
			3003: "Tablet",
		}

		// Find by product ID
		got, found := maps.FirstEntryBy(m, func(p productID) int { return p.id },
			productID{2002, "iPhone"},
			productID{3003, "iPad"},
		)
		if !found || got.Key.name != "iPhone" || got.Val != "Phone" {
			t.Errorf("FirstEntryBy() with custom key type = %+v, want key={2002 iPhone} value=Phone", got)
		}

		// Not found case
		_, found = maps.FirstEntryBy(m, func(p productID) int { return p.id }, productID{9999, "Unknown"})
		if found {
			t.Error("Expected not found for non-existent key")
		}
	})
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
				stringGotKey, stringGotFound := maps.FirstKeyBy(tt.m, trans, k...)
				gotKey = stringGotKey
				gotFound = stringGotFound
			case []int:
				trans, ok := tt.transform.(func(int) string)
				if !ok {
					t.Fatalf("Test case %q: expected transform to be func(int) string, got %T", tt.name, tt.transform)
				}
				// Call FirstKeyBy with int keys and int transform (transform is now the second argument)
				intGotKey, intGotFound := maps.FirstKeyBy(tt.m, trans, k...)
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

func TestFirstKeyOrRandom(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	t.Run("key found", func(t *testing.T) {
		// The behavior is deterministic here.
		key := maps.FirstKeyOrRandom(m, "x", "y", "b", "a")
		assert.Equal(t, "b", key)
	})

	t.Run("key not found, fallback to random", func(t *testing.T) {
		key := maps.FirstKeyOrRandom(m, "x", "y", "z")
		// The behavior is random, so we check if the key is valid.
		_, exists := m[key]
		assert.True(t, exists, "fallback key should be a valid key from the map")
	})

	t.Run("empty map", func(t *testing.T) {
		emptyMap := map[string]int{}
		key := maps.FirstKeyOrRandom(emptyMap, "x", "y")
		assert.Zero(t, key)
	})
}

func TestFirstValueOrRandom(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	t.Run("key found", func(t *testing.T) {
		// The behavior is deterministic here.
		value := maps.FirstValueOrRandom(m, "x", "y", "c", "a")
		assert.Equal(t, 3, value)
	})

	t.Run("key not found, fallback to random", func(t *testing.T) {
		value := maps.FirstValueOrRandom(m, "x", "y", "z")
		// The behavior is random, so we check if the value is valid.
		found := false
		for _, v := range m {
			if v == value {
				found = true
				break
			}
		}
		assert.True(t, found, "fallback value should be a valid value from the map")
	})

	t.Run("empty map", func(t *testing.T) {
		emptyMap := map[string]int{}
		value := maps.FirstValueOrRandom(emptyMap, "x", "y")
		assert.Zero(t, value)
	})
}
