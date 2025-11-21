//go:build !race

package maps_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goexts/generic/maps"
)

func TestToKVs(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]string
		want []maps.KeyValue[int, string]
	}{
		{
			name: "empty map",
			m:    map[int]string{},
			want: maps.KVs[int, string](),
		},
		{
			name: "single element",
			m:    map[int]string{1: "a"},
			want: maps.KVs(maps.KV(1, "a")),
		},
		{
			name: "multiple elements",
			m:    map[int]string{1: "a", 2: "b", 3: "c"},
			want: maps.KVs(
				maps.KV(1, "a"),
				maps.KV(2, "b"),
				maps.KV(3, "c"),
			),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maps.ToKVs(tt.m)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestFromKVs(t *testing.T) {
	tests := []struct {
		name string
		kvs  []maps.KeyValue[string, int]
		want map[string]int
	}{
		{
			name: "empty slice",
			kvs:  maps.KVs[string, int](),
			want: map[string]int{},
		},
		{
			name: "single element",
			kvs:  maps.KVs(maps.KV("a", 1)),
			want: map[string]int{"a": 1},
		},
		{
			name: "multiple elements",
			kvs: maps.KVs(
				maps.KV("a", 1),
				maps.KV("b", 2),
			),
			want: map[string]int{"a": 1, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// By assigning the result to a declared variable, we allow the compiler
			// to infer the generic types K, V, and M.
			var got = maps.FromKVs(tt.kvs...)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestToSliceWith(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		f    func(string, int) (string, bool)
		want []string
	}{
		{
			name: "empty map",
			m:    map[string]int{},
			f:    func(k string, v int) (string, bool) { return fmt.Sprintf("%s:%d", k, v), true },
			want: []string{},
		},
		{
			name: "filter out all elements",
			m:    map[string]int{"a": 1, "b": 2, "c": 3},
			f:    func(_ string, _ int) (string, bool) { return "", false },
			want: []string{},
		},
		{
			name: "filter some elements",
			m:    map[string]int{"a": 1, "b": 2, "c": 3, "d": 4},
			f: func(k string, v int) (string, bool) {
				if v%2 == 0 {
					return fmt.Sprintf("%s:%d", k, v), true
				}
				return "", false
			},
			want: []string{"b:2", "d:4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maps.ToSliceWith(tt.m, tt.f)
			assert.ElementsMatch(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		var fn func(string, int) (string, bool)
		assert.Panics(t, func() {
			maps.ToSliceWith(map[string]int{"a": 1}, fn)
		})
	})
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		f    func(string, int) string
		want []string
	}{
		{
			name: "empty map",
			m:    map[string]int{},
			f:    func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) },
			want: []string{},
		},
		{
			name: "single element",
			m:    map[string]int{"a": 1},
			f:    func(k string, v int) string { return fmt.Sprintf("%s:%d", k, v) },
			want: []string{"a:1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maps.ToSlice(tt.m, tt.f)
			assert.ElementsMatch(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		var fn func(string, int) string
		assert.Panics(t, func() {
			maps.ToSlice(map[string]int{"a": 1}, fn)
		})
	})
}

func TestFromSliceWithIndex(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		f    func(int, string) (string, int)
		want map[string]int
	}{
		{
			name: "empty slice",
			s:    []string{},
			f: func(i int, s string) (string, int) {
				return fmt.Sprintf("%s_%d", s, i), len(s)
			},
			want: map[string]int{},
		},
		{
			name: "single element",
			s:    []string{"a"},
			f: func(i int, s string) (string, int) {
				return fmt.Sprintf("%s_%d", s, i), len(s) * (i + 1)
			},
			want: map[string]int{"a_0": 1},
		},
		{
			name: "multiple elements with index",
			s:    []string{"a", "bb", "ccc"},
			f: func(i int, s string) (string, int) {
				return fmt.Sprintf("%s_%d", s, i), len(s) * (i + 1)
			},
			want: map[string]int{
				"a_0":   1, // 1 * (0 + 1)
				"bb_1":  4, // 2 * (1 + 1)
				"ccc_2": 9, // 3 * (2 + 1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = maps.FromSliceWithIndex(tt.s, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		var fn func(int, any) (string, int)
		assert.Panics(t, func() {
			maps.FromSliceWithIndex([]any{"test"}, fn)
		})
	})
}

func TestFromSlice(t *testing.T) {
	tests := []struct {
		name string
		s    []string
		f    func(string) (string, int)
		want map[string]int
	}{
		{
			name: "empty slice",
			s:    []string{},
			f: func(s string) (string, int) {
				return s, len(s)
			},
			want: map[string]int{},
		},
		{
			name: "single element",
			s:    []string{"a"},
			f: func(s string) (string, int) {
				return s, len(s)
			},
			want: map[string]int{"a": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got = maps.FromSlice(tt.s, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		var fn func(string) (string, int)
		assert.Panics(t, func() {
			maps.FromSlice([]string{"a"}, fn)
		})
	})
}

func TestFromSlice_ComplexType(t *testing.T) {
	type User struct {
		ID int
	}
	users := []*User{
		{ID: 1},
		{ID: 2},
	}

	// By assigning the result to a declared variable, we allow the compiler
	// to infer the generic types K, V, and M.
	var got = maps.FromSlice(users, func(u *User) (int, *User) {
		return u.ID, u
	})

	expected := map[int]*User{
		1: users[0],
		2: users[1],
	}
	assert.Equal(t, expected, got)
}

func TestMerge(t *testing.T) {
	t.Run("merge with overlay true", func(t *testing.T) {
		dest := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}
		maps.Merge(dest, src, true)
		assert.Equal(t, map[string]int{"a": 1, "b": 3, "c": 4}, dest)
	})

	t.Run("merge with overlay false", func(t *testing.T) {
		dest := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}
		maps.Merge(dest, src, false)
		assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 4}, dest)
	})
}

func TestMergeWith(t *testing.T) {
	t.Run("merge with custom function", func(t *testing.T) {
		dest := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}
		maps.MergeWith(dest, src, func(_ string, oldV, newV int) int {
			return oldV + newV
		})
		assert.Equal(t, map[string]int{"a": 1, "b": 5, "c": 4}, dest)
	})
}

func TestConcat(t *testing.T) {
	t.Run("concat multiple maps", func(t *testing.T) {
		dest := map[string]int{"a": 1}
		m1 := map[string]int{"b": 2}
		m2 := map[string]int{"c": 3}
		maps.Concat(dest, m1, m2)
		assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, dest)
	})
}

func TestConcatWith(t *testing.T) {
	t.Run("concat with custom function", func(t *testing.T) {
		dest := map[string]int{"a": 1}
		m1 := map[string]int{"a": 2, "b": 3}
		m2 := map[string]int{"b": 4, "c": 5}
		maps.ConcatWith(func(_ string, v1, v2 int) int {
			return v1 + v2
		}, dest, m1, m2)
		assert.Equal(t, map[string]int{"a": 3, "b": 7, "c": 5}, dest)
	})
}

func TestExclude(t *testing.T) {
	t.Run("exclude keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		maps.Exclude(m, "a", "c")
		assert.Equal(t, map[string]int{"b": 2}, m)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter even values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		maps.Filter(m, func(_ string, v int) bool {
			return v%2 == 0
		})
		assert.Equal(t, map[string]int{"b": 2, "d": 4}, m)
	})
}

func TestToStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	tests := []struct {
		name string
		m    map[string]interface{}
		f    func(*Person, string, interface{}) *Person
		want *Person
	}{
		{
			name: "empty map",
			m:    map[string]interface{}{},
			f: func(p *Person, k string, v interface{}) *Person {
				if p == nil {
					p = &Person{}
				}
				switch k {
				case "name":
					p.Name = v.(string)
				case "age":
					p.Age = v.(int)
				}
				return p
			},
			want: &Person{},
		},
		{
			name: "with values",
			m: map[string]interface{}{
				"name": "John",
				"age":  30,
			},
			f: func(p *Person, k string, v interface{}) *Person {
				if p == nil {
					p = &Person{}
				}
				switch k {
				case "name":
					p.Name = v.(string)
				case "age":
					p.Age = v.(int)
				}
				return p
			},
			want: &Person{Name: "John", Age: 30},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maps.ToStruct(tt.m, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTransform(t *testing.T) {
	t.Run("transform keys and values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := maps.Transform(m, func(k string, v int) (string, int, bool) {
			return strings.ToUpper(k), v * 2, true
		})
		assert.Equal(t, map[string]int{"A": 2, "B": 4}, result)
	})

	t.Run("skip some items", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		result := maps.Transform(m, func(k string, v int) (string, int, bool) {
			if v%2 == 0 {
				return "", 0, false
			}
			return k, v, true
		})
		assert.Equal(t, map[string]int{"a": 1, "c": 3}, result)
	})
}

func TestRandom(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := map[int]string{}
		k, v, ok := maps.Random(m)
		assert.False(t, ok)
		assert.Zero(t, k)
		assert.Zero(t, v)
	})

	t.Run("non-empty map", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		k, v, ok := maps.Random(m)
		require.True(t, ok)

		// Assert that the returned key-value pair is valid.
		expectedValue, keyExists := m[k]
		assert.True(t, keyExists, "random key should exist in the map")
		assert.Equal(t, expectedValue, v, "value should match the key")
	})
}

func TestRandomKey(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := map[int]string{}
		k, ok := maps.RandomKey(m)
		assert.False(t, ok)
		assert.Zero(t, k)
	})

	t.Run("non-empty map", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		k, ok := maps.RandomKey(m)
		require.True(t, ok)

		// Assert that the returned key exists in the map.
		_, keyExists := m[k]
		assert.True(t, keyExists, "random key should exist in the map")
	})
}

func TestRandomValue(t *testing.T) {
	t.Run("empty map", func(t *testing.T) {
		m := map[int]string{}
		v, ok := maps.RandomValue(m)
		assert.False(t, ok)
		assert.Zero(t, v)
	})

	t.Run("non-empty map", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		randomVal, ok := maps.RandomValue(m)
		require.True(t, ok)

		// Assert that the returned value exists in the map's values.
		found := false
		for _, v := range m {
			if v == randomVal {
				found = true
				break
			}
		}
		assert.True(t, found, "random value should be one of the map's values")
	})
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
