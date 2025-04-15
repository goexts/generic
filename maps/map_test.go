//go:build !race

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package maps implements the functions, types, and interfaces for the module.
package maps

import (
	"fmt"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func sortKV(kvs []KeyValue[int, string]) []KeyValue[int, string] {
	slices.SortFunc(kvs, func(a, b KeyValue[int, string]) int {
		if a.Key > b.Key {
			return 1
		} else if a.Key < b.Key {
			return -1
		}
		return 0
	})
	return kvs
}
func TestMapToKVs(t *testing.T) {
	type args struct {
		m map[int]string
	}
	type testCase struct {
		name string
		args args
		want []KeyValue[int, string]
	}
	tests := []testCase{
		{
			name: "test",
			args: args{
				m: map[int]string{
					1: "a",
					2: "b",
					3: "c",
				},
			},
			want: []KeyValue[int, string]{
				{
					Key: 3,
					Val: "c",
				},
				{
					Key: 2,
					Val: "b",
				},
				{
					Key: 1,
					Val: "a",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapToKVs(tt.args.m)
			assert.EqualValues(t, sortKV(tt.want), sortKV(got))
		})
	}
}

// Merging multiple non-empty maps with unique keys
func TestMergeMapsWithUniqueKeys(t *testing.T) {
	// Arr matey! Let's merge these treasure maps!

	result := map[string]int{"gold": 100, "gems": 50}
	map2 := map[string]int{"coins": 75, "pearls": 25}

	MergeMaps(result, map2)

	assert.Equal(t, 4, len(result))
	assert.Equal(t, 100, result["gold"])
	assert.Equal(t, 50, result["gems"])
	assert.Equal(t, 75, result["coins"])
	assert.Equal(t, 25, result["pearls"])
}

// Merging two maps with overlapping keys where second map values override first
func TestMergeMapsWithOverlappingKeys(t *testing.T) {
	// Shiver me timbers! These maps be fighting over the same booty!

	result := map[string]int{"chest": 100, "map": 1}
	map2 := map[string]int{"chest": 200, "compass": 1}

	MergeMaps(result, map2)

	assert.Equal(t, 3, len(result))
	assert.Equal(t, 200, result["chest"])
	assert.Equal(t, 1, result["map"])
	assert.Equal(t, 1, result["compass"])
}

// Merging three or more maps in sequence
func TestMergeMapsSequence(t *testing.T) {
	// Yo ho ho! A fleet of maps approaching!

	result := map[string]int{"a": 1}
	map2 := map[string]int{"b": 2}
	map3 := map[string]int{"c": 3}
	map4 := map[string]int{"d": 4}

	MergeMaps(result, map2, map3, map4)

	assert.Equal(t, 4, len(result))
	assert.Equal(t, 1, result["a"])
	assert.Equal(t, 2, result["b"])
	assert.Equal(t, 3, result["c"])
	assert.Equal(t, 4, result["d"])
}

// Passing empty variadic maps argument returns nil
func TestMergeMapsEmptyVariadic(t *testing.T) {
	// Blimey! The map chest be empty!

	base := map[string]int{"x": 1}
	MergeMaps(base)

	assert.NotNil(t, base)
}

// Merging nil map with non-nil map
func TestMergeMapsWithNilMap(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Walk the plank! Expected panic for nil map but did not get one")
		}
	}()
	var nilMap map[string]int
	validMap := map[string]int{"valid": 1}

	MergeMaps(nilMap, validMap)

	assert.Fail(t, "panic expected")
}

// Merging maps with zero capacity
func TestMergeMapsZeroCapacity(t *testing.T) {
	// Dead men's maps tell no tales!

	result := make(map[string]int, 0)
	map2 := make(map[string]int, 0)

	result["key"] = 1
	map2["another"] = 2

	MergeMaps(result, map2)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, 1, result["key"])
	assert.Equal(t, 2, result["another"])
}

// Merging maps containing nil values
func TestMergeMapsWithNilValues(t *testing.T) {
	// Ghost values in the map! Spooky!

	result := map[string]*string{"key1": nil}
	map2 := map[string]*string{"key2": nil}

	MergeMaps(result, map2)

	assert.Equal(t, 2, len(result))
	assert.Nil(t, result["key1"])
	assert.Nil(t, result["key2"])
}

// Map empty input map to empty slice
func TestMapToTypesEmptyMap(t *testing.T) {
	// Arr! Empty maps be like empty treasure chests - disappointing but valid!

	emptyMap := make(map[string]int)

	result := MapToTypes(emptyMap, func(k string, v int) string {
		return fmt.Sprintf("%s:%d", k, v)
	})

	assert.Empty(t, result)
}

// Map single key-value pair to single element slice
func TestMapToTypesSinglePair(t *testing.T) {
	// Yo ho ho! A lone key-value pair, like a single gold coin in me pocket!

	m := map[string]int{"key": 42}

	result := MapToTypes(m, func(k string, v int) string {
		return fmt.Sprintf("%s:%d", k, v)
	})

	assert.Equal(t, []string{"key:42"}, result)
}

// Map multiple key-value pairs to slice preserving order
func TestMapToTypesMultiplePairs(t *testing.T) {
	// Shiver me timbers! A whole chest of key-value booty!

	m := map[string]int{
		"first":  1,
		"second": 2,
		"third":  3,
	}

	result := MapToTypes(m, func(k string, v int) string {
		return fmt.Sprintf("%s:%d", k, v)
	})

	assert.Len(t, result, 3)
	assert.Contains(t, result, "first:1")
	assert.Contains(t, result, "second:2")
	assert.Contains(t, result, "third:3")
}

// Handle nil map input
func TestMapToTypesNilMap(t *testing.T) {
	// Blimey! A nil map be like a sunken ship - nothing good comes from it!

	var nilMap map[string]int

	result := MapToTypes(nilMap, func(k string, v int) string {
		return fmt.Sprintf("%s:%d", k, v)
	})

	assert.Empty(t, result)
}

// Handle nil transform function
func TestMapToTypesNilFunction(t *testing.T) {
	// Arrr! Passing nil function be asking for trouble, like sailing without a compass!
	m := map[string]int{"key": 42}

	assert.Panics(t, func() {
		MapToTypes[map[string]int, string, int, string](m, nil)
	})
}

// Process map with zero capacity
func TestMapToTypesZeroCapMap(t *testing.T) {
	// Dead men tell no tales, and zero cap maps hold no values!

	m := make(map[string]int)
	m["key"] = 42

	result := MapToTypes(m, func(k string, v int) string {
		return fmt.Sprintf("%s:%d", k, v)
	})

	assert.Equal(t, []string{"key:42"}, result)
}

// Handle maps with interface{} values
func TestMapToTypesInterfaceValues(t *testing.T) {
	// By Blackbeard's beard! These interface values be shape-shifters!
	m := map[string]interface{}{
		"str":  "hello",
		"num":  42,
		"bool": true,
	}

	result := MapToTypes(m, func(k string, v interface{}) string {
		return fmt.Sprintf("%v", v)
	})

	assert.Len(t, result, 3)
	assert.Contains(t, result, "hello")
	assert.Contains(t, result, "42")
	assert.Contains(t, result, "true")
}

type MapToTypesStruct struct {
	Str  string
	Num  int
	Bool bool
}

// Handle maps with interface{} values
func TestMapToTypesStruct(t *testing.T) {
	// By Blackbeard's beard! These interface values be shape-shifters!
	m := map[string]interface{}{
		"str":  "hello",
		"num":  42,
		"bool": true,
	}

	result := MapToTypes(m, func(k string, v interface{}) string {
		return fmt.Sprintf("%v", v)
	})

	assert.Len(t, result, 3)
	assert.Contains(t, result, "hello")
	assert.Contains(t, result, "42")
	assert.Contains(t, result, "true")
}

func TestRemap(t *testing.T) {
	src := map[string]int{
		"a": 1,
		"b": 2,
	}
	s := Transform(src, func(k string, v int) (string, any, bool) {
		return k, v, true
	})
	t.Logf("%+v", s)
}
