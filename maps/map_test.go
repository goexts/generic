//go:build !race

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package maps implements the functions, types, and interfaces for the module.
package maps

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToKVs(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]string
		want []KeyValue[int, string]
	}{
		{
			name: "empty map",
			m:    map[int]string{},
			want: []KeyValue[int, string]{},
		},
		{
			name: "single element",
			m:    map[int]string{1: "a"},
			want: []KeyValue[int, string]{{Key: 1, Val: "a"}},
		},
		{
			name: "multiple elements",
			m:    map[int]string{1: "a", 2: "b", 3: "c"},
			want: []KeyValue[int, string]{
				{Key: 1, Val: "a"},
				{Key: 2, Val: "b"},
				{Key: 3, Val: "c"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ToKVs(tt.m)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestFromKVs(t *testing.T) {
	tests := []struct {
		name string
		kvs  []KeyValue[string, int]
		want map[string]int
	}{
		{
			name: "empty slice",
			kvs:  []KeyValue[string, int]{},
			want: map[string]int{},
		},
		{
			name: "single element",
			kvs:  []KeyValue[string, int]{{Key: "a", Val: 1}},
			want: map[string]int{"a": 1},
		},
		{
			name: "multiple elements",
			kvs: []KeyValue[string, int]{
				{Key: "a", Val: 1},
				{Key: "b", Val: 2},
			},
			want: map[string]int{"a": 1, "b": 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FromKVs[string, int, map[string]int](tt.kvs)
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
			got := ToSliceWith(tt.m, tt.f)
			assert.ElementsMatch(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		assert.Panics(t, func() {
			ToSliceWith[map[string]int, string, int, string](map[string]int{"a": 1}, nil)
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
			got := ToSlice(tt.m, tt.f)
			assert.ElementsMatch(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		assert.Panics(t, func() {
			ToSlice[map[string]int, string, int, string](map[string]int{"a": 1}, nil)
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
			got := FromSliceWithIndex[string, map[string]int, string, int](tt.s, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		assert.Panics(t, func() {
			FromSliceWithIndex[any, map[string]int, string, int]([]any{"test"}, nil)
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
			got := FromSlice[string, map[string]int, string, int](tt.s, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}

	// Test with nil function
	t.Run("nil function", func(t *testing.T) {
		assert.Panics(t, func() {
			FromSlice[string, map[string]int, string, int]([]string{"a"}, nil)
		})
	})
}

func TestMerge(t *testing.T) {
	t.Run("merge with overlay true", func(t *testing.T) {
		dest := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}
		Merge(dest, src, true)
		assert.Equal(t, map[string]int{"a": 1, "b": 3, "c": 4}, dest)
	})

	t.Run("merge with overlay false", func(t *testing.T) {
		dest := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}
		Merge(dest, src, false)
		assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 4}, dest)
	})
}

func TestMergeWith(t *testing.T) {
	t.Run("merge with custom function", func(t *testing.T) {
		dest := map[string]int{"a": 1, "b": 2}
		src := map[string]int{"b": 3, "c": 4}
		MergeWith(dest, src, func(_ string, oldV, newV int) int {
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
		Concat(dest, m1, m2)
		assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, dest)
	})
}

func TestConcatWith(t *testing.T) {
	t.Run("concat with custom function", func(t *testing.T) {
		dest := map[string]int{"a": 1}
		m1 := map[string]int{"a": 2, "b": 3}
		m2 := map[string]int{"b": 4, "c": 5}
		ConcatWith(func(_ string, v1, v2 int) int {
			return v1 + v2
		}, dest, m1, m2)
		assert.Equal(t, map[string]int{"a": 3, "b": 7, "c": 5}, dest)
	})
}

func TestExclude(t *testing.T) {
	t.Run("exclude keys", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		Exclude(m, "a", "c")
		assert.Equal(t, map[string]int{"b": 2}, m)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter even values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
		Filter(m, func(_ string, v int) bool {
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
			got := ToStruct(tt.m, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTransform(t *testing.T) {
	t.Run("transform keys and values", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2}
		result := Transform(m, func(k string, v int) (string, int, bool) {
			return strings.ToUpper(k), v * 2, true
		})
		assert.Equal(t, map[string]int{"A": 2, "B": 4}, result)
	})

	t.Run("skip some items", func(t *testing.T) {
		m := map[string]int{"a": 1, "b": 2, "c": 3}
		result := Transform(m, func(k string, v int) (string, int, bool) {
			if v%2 == 0 {
				return "", 0, false
			}
			return k, v, true
		})
		assert.Equal(t, map[string]int{"a": 1, "c": 3}, result)
	})
}
