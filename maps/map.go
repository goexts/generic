package maps

import (
	"golang.org/x/exp/maps"
)

// Keys returns the keys of the map m.
// The keys will be in an indeterminate order.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return maps.Keys(m)
}

// Values returns the values of the map m.
// The values will be in an indeterminate order.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return maps.Values(m)
}

// Equal reports whether two maps contain the same key/value pairs.
// Values are compared using ==.
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	return maps.Equal(m1, m2)
}

// EqualFunc is like Equal, but compares values using eq.
// Keys are still compared with ==.
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool {
	return maps.EqualFunc(m1, m2, eq)
}

// Clear removes all entries from m, leaving it empty.
func Clear[M ~map[K]V, K comparable, V any](m M) {
	maps.Clear(m)
}

// Clone returns a copy of m.  This is a shallow clone:
// the new keys and values are set using ordinary assignment.
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	return maps.Clone(m)
}

// Copy copies all key/value pairs in src adding them to dst.
// When a key in src is already present in dst,
// the value in dst will be overwritten by the value associated
// with the key in src.
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2) {
	maps.Copy(dst, src)
}

// DeleteFunc deletes any key/value pairs from m for which del returns true.
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool) {
	maps.DeleteFunc(m, del)
}

// Merge merges the values of src into dest.
func Merge[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dest M1, src M2, overlay bool) {
	for k, v := range src {
		if _, ok := dest[k]; !ok || overlay {
			dest[k] = v
		}
	}
}

// MergeFunc merges the values of src into dest using the provided merge function.
func MergeFunc[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dest M1, src M2, merge func(K, V, V) V) {
	for k, v := range src {
		if _, ok := dest[k]; !ok {
			dest[k] = v
		} else {
			dest[k] = merge(k, dest[k], v)
		}
	}
}

// Filter removes all key/value pairs from m for which f returns false.
func Filter[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) {
	for k, v := range m {
		if !f(k, v) {
			delete(m, k)
		}
	}
}

// FilterFunc is like Filter, but uses a function.
func FilterFunc[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) {
	for k, v := range m {
		if !f(k, v) {
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
