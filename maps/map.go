package maps

// Keys returns a slice of the keys in the map.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	var ks []K
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

// Values returns a slice of the values in the map.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	var vs []V
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

// KeyValue is a key-value pair.
type KeyValue[K comparable, V any] struct {
	Key K
	Val V
}

// MapToKVs converts a map to a slice of key-value pairs.
func MapToKVs[M ~map[K]V, K comparable, V any, KV KeyValue[K, V]](m M) []KV {
	var kvs []KV
	for k, v := range m {
		kvs = append(kvs, KV{Key: k, Val: v})
	}
	return kvs
}

// MapToTypes converts a map to a slice of types.
func MapToTypes[M ~map[K]V, K comparable, V any, T any](m M, f func(K, V) T) []T {
	var ts []T
	for k, v := range m {
		ts = append(ts, f(k, v))
	}
	return ts
}
