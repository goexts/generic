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

// MergeFunc merges the values of src into dest using the provided merge function.
// If a key exists in both maps, the merge function will be called to determine the final value.
func MergeFunc[M ~map[K]V, K comparable, V any](dest M, src M, cmp func(key K, src V, val V) V) {
	for k, v := range src {
		if existing, ok := dest[k]; !ok {
			dest[k] = v
		} else {
			dest[k] = cmp(k, existing, v)
		}
	}
}

// MergeMaps merges multiple maps into a single map.
// If a key exists in multiple maps, the value from the last map will be used.
func MergeMaps[M ~map[K]V, K comparable, V any](m M, ms ...M) {
	if len(ms) == 0 {
		return
	}

	// Merge subsequent ms into the result map
	for _, mm := range ms {
		Merge(m, mm, true)
	}
}

// MergeMapsFunc merges multiple maps into a single map using a custom merge function.
// If a key exists in multiple maps, the merge function will be called to determine the final value.
func MergeMapsFunc[M ~map[K]V, K comparable, V any](merge func(K, V, V) V, m M, ms ...M) {
	if len(ms) == 0 {
		return
	}

	// Merge subsequent maps into the result map
	for _, mm := range ms {
		MergeFunc(m, mm, merge)
	}
}

// Filter removes all key/value pairs from m for which f returns false.
func Filter[M ~map[K]V, K comparable, V any](m M, keys ...K) {
	for i := range keys {
		if _, ok := m[keys[i]]; ok {
			delete(m, keys[i])
		}
	}
}

// FilterFunc is like Filter, but uses a function.
func FilterFunc[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) {
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

// MapToKVs converts a map to a slice of key-value pairs.
func MapToKVs[M ~map[K]V, K comparable, V any, KV KeyValue[K, V]](m M) []KV {
	kvs := make([]KV, 0, len(m))
	for k, v := range m {
		kvs = append(kvs, KV{Key: k, Val: v})
	}
	return kvs
}

// KVsToMap converts a slice of key-value pairs to a map.
func KVsToMap[KV KeyValue[K, V], K comparable, V any, M ~map[K]V](kvs []KeyValue[K, V]) M {
	m := make(M, len(kvs))
	for _, kv := range kvs {
		m[kv.Key] = kv.Val
	}
	return m
}

// MapToTypes converts a map to a slice of types.
func MapToTypes[M ~map[K]V, K comparable, V any, T any](m M, f func(K, V) T) []T {
	ts := make([]T, 0, len(m))
	for k, v := range m {
		ts = append(ts, f(k, v))
	}
	return ts
}

// TypesToMap converts a slice of types to a map.
func TypesToMap[T any, M ~map[K]V, K comparable, V any](ts []T, f func(T) (K, V)) M {
	m := make(M, len(ts))
	for _, t := range ts {
		k, v := f(t)
		m[k] = v
	}
	return m
}

// MapToStruct converts a map to a struct.
func MapToStruct[M ~map[K]V, K comparable, V any, S any](m M, f func(*S, K, V) *S) *S {
	s := new(S)
	for k, v := range m {
		s = f(s, k, v)
	}
	return s
}
