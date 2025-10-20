/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package maps implements the functions, types, and interfaces for the module.
package maps

// Merge merges the values of src into dest.
// If overlay is true, existing values in dest will be overwritten.
func Merge[M ~map[K]V, K comparable, V any](dest M, src M, overlay bool) {
	for k, v := range src {
		if _, ok := dest[k]; !ok || overlay {
			dest[k] = v
		}
	}
}

// MergeWith merges the values of src into dest using the provided merge function.
// If a key exists in both maps, the merge function will be called to determine the final value.
func MergeWith[M ~map[K]V, K comparable, V any](dest M, src M, cmp func(key K, src V, val V) V) {
	for k, v := range src {
		if existing, ok := dest[k]; ok {
			dest[k] = cmp(k, existing, v)
		} else {
			dest[k] = v
		}
	}
}

// Concat merges multiple maps into a single map.
// If a key exists in multiple maps, the value from the last map will be used.
func Concat[M ~map[K]V, K comparable, V any](m M, ms ...M) {
	if len(ms) == 0 {
		return
	}

	// Merge subsequent ms into the result map
	for _, mm := range ms {
		Merge(m, mm, true)
	}
}

// ConcatWith merges multiple maps into a single map using a custom merge function.
// If a key exists in multiple maps, the merge function will be called to determine the final value.
func ConcatWith[M ~map[K]V, K comparable, V any](merge func(K, V, V) V, m M, ms ...M) {
	if len(ms) == 0 {
		return
	}

	// Merge subsequent maps into the result map
	for _, mm := range ms {
		MergeWith(m, mm, merge)
	}
}

// Exclude removes all key/value pairs from m for which f returns false.
func Exclude[M ~map[K]V, K comparable, V any](m M, keys ...K) {
	for i := range keys {
		delete(m, keys[i])
	}
}

// Filter is like Exclude, but uses a function.
func Filter[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) {
	for k, v := range m {
		if f(k, v) {
			delete(m, k)
		}
	}
}

// KeyValue is a key-value pair.
type KeyValue[K comparable, V any] struct {
	Key K
	Val V
}

// ToKVs converts a map to a slice of key-value pairs.
func ToKVs[M ~map[K]V, K comparable, V any](m M) []KeyValue[K, V] {
	kvs := make([]KeyValue[K, V], 0, len(m))
	for k, v := range m {
		kvs = append(kvs, KeyValue[K, V]{Key: k, Val: v})
	}
	return kvs
}

// FromKVs converts a slice of key-value pairs to a map.
func FromKVs[K comparable, V any, M ~map[K]V](kvs []KeyValue[K, V]) M {
	m := make(M, len(kvs))
	for _, kv := range kvs {
		m[kv.Key] = kv.Val
	}
	return m
}

// ToSlice converts a map to a slice of types.
func ToSlice[M ~map[K]V, K comparable, V any, T any](m M, f func(K, V) T) []T {
	ts := make([]T, 0, len(m))
	for k, v := range m {
		ts = append(ts, f(k, v))
	}
	return ts
}

// FromSlice converts a slice of types to a map.
func FromSlice[T any, M ~map[K]V, K comparable, V any](ts []T, f func(T) (K, V)) M {
	m := make(M, len(ts))
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}

// ToStruct converts a map to a struct.
func ToStruct[M ~map[K]V, K comparable, V any, S any](m M, f func(*S, K, V) *S) *S {
	s := new(S)
	for k, v := range m {
		s = f(s, k, v)
	}
	return s
}

// Transform remaps the keys and values of a map using a custom transformation function.
// The transformation function is called for each key-value pair in the original map.
// If the transformation function returns false as its third return value, the key-value pair is skipped.
// Otherwise, the transformed key-value pair is added to the new map.
func Transform[M ~map[K]V, K comparable, V any, TK comparable, TV any](m M, f func(K, V) (TK, TV, bool)) map[TK]TV {
	// Create a new map with the same length as the original map to avoid reallocations.
	n := make(map[TK]TV, len(m))

	// Iterate over each key-value pair in the original map.
	for k, v := range m {
		// Call the transformation function to get the new key, value, and a boolean indicating whether to include the pair.
		j, w, ok := f(k, v)

		// If the transformation function returned false, skip this key-value pair.
		if !ok {
			continue
		}

		// Add the transformed key-value pair to the new map.
		n[j] = w
	}

	// Return the new map with the transformed key-value pairs.
	return n
}
