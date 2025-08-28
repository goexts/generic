package types

import (
	"bytes"
	"reflect"

	"github.com/goexts/generic/object"
)

// isNil checks if the provided value is nil. It's used for generic types
// that are constrained by interfaces, as they cannot be directly compared to nil.
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.UnsafePointer, reflect.Interface, reflect.Slice:
		return rv.IsNil()
	default:
		return false
	}
}

// entry represents a key-value pair in the Map.
type entry[K object.Object, V object.Object] struct {
	key   K
	value V
}

// Map is a wrapper that behaves like a hash map, using the object's
// HashCode and Equals methods for key management. It is designed to be
// similar to java.util.Map.
type Map[K object.Object, V object.Object] struct {
	object.BaseObject
	data map[int][]*entry[K, V]
	size int
}

// NewMap creates a new empty Map.
func NewMap[K object.Object, V object.Object]() *Map[K, V] {
	return &Map[K, V]{
		data: make(map[int][]*entry[K, V]),
		size: 0,
	}
}

// findEntry finds the entry for a given key. It returns the bucket and the index within the bucket.
// If not found, the index is -1.
func (m *Map[K, V]) findEntry(key K) ([]*entry[K, V], int) {
	if isNil(key) { // Handle nil key
		bucket := m.data[0]
		for i, e := range bucket {
			if isNil(e.key) {
				return bucket, i
			}
		}
		return bucket, -1
	}

	hash := key.HashCode()
	bucket := m.data[hash]
	for i, e := range bucket {
		if !isNil(e.key) && e.key.Equals(key) {
			return bucket, i
		}
	}
	return bucket, -1
}

// Put associates the specified value with the specified key in this map.
// It returns the previous value associated with the key, or the zero value of V
// if there was no mapping for the key.
func (m *Map[K, V]) Put(key K, value V) V {
	bucket, index := m.findEntry(key)

	if index != -1 { // Key found, update value
		oldValue := bucket[index].value
		bucket[index].value = value
		return oldValue
	}

	// Key not found, add new entry
	hash := 0
	if !isNil(key) {
		hash = key.HashCode()
	}
	m.data[hash] = append(m.data[hash], &entry[K, V]{key: key, value: value})
	m.size++
	var zeroV V
	return zeroV
}

// Get returns the value to which the specified key is mapped.
func (m *Map[K, V]) Get(key K) (V, bool) {
	bucket, index := m.findEntry(key)
	if index != -1 {
		return bucket[index].value, true
	}
	var zeroV V
	return zeroV, false
}

// Remove removes the mapping for a key from this map if it is present.
// It returns the value that was removed, or the zero value of V if the key was not found.
func (m *Map[K, V]) Remove(key K) V {
	hash := 0
	if !isNil(key) {
		hash = key.HashCode()
	}

	bucket, index := m.findEntry(key)
	if index != -1 {
		oldValue := bucket[index].value
		// Remove element from slice by swapping with the last element
		bucket[index] = bucket[len(bucket)-1]
		m.data[hash] = bucket[:len(bucket)-1]
		m.size--
		return oldValue
	}
	var zeroV V
	return zeroV
}

// Size returns the number of key-value mappings in this map.
func (m *Map[K, V]) Size() int {
	return m.size
}

// IsEmpty returns true if this map contains no key-value mappings.
func (m *Map[K, V]) IsEmpty() bool {
	return m.size == 0
}

// ContainsKey returns true if this map contains a mapping for the specified key.
func (m *Map[K, V]) ContainsKey(key K) bool {
	_, found := m.Get(key)
	return found
}

// Keys returns a slice containing all of the keys in this map.
func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, m.size)
	for _, bucket := range m.data {
		for _, e := range bucket {
			keys = append(keys, e.key)
		}
	}
	return keys
}

// Values returns a slice containing all of the values in this map.
func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, m.size)
	for _, bucket := range m.data {
		for _, e := range bucket {
			values = append(values, e.value)
		}
	}
	return values
}

// String returns a string representation of the map.
func (m *Map[K, V]) String() string {
	var buf bytes.Buffer
	buf.WriteString("{")
	first := true
	for _, bucket := range m.data {
		for _, e := range bucket {
			if !first {
				buf.WriteString(", ")
			}
			if isNil(e.key) {
				buf.WriteString("null")
			} else {
				buf.WriteString(e.key.String())
			}
			buf.WriteString(": ")
			if isNil(e.value) {
				buf.WriteString("null")
			} else {
				buf.WriteString(e.value.String())
			}
			first = false
		}
	}
	buf.WriteString("}")
	return buf.String()
}

// Equals checks if two Map objects are equal.
func (m *Map[K, V]) Equals(other object.Object) bool {
	if isNil(other) {
		return false
	}
	m2, ok := other.(*Map[K, V])
	if !ok || m.Size() != m2.Size() {
		return false
	}

	for _, bucket := range m.data {
		for _, e1 := range bucket {
			v2, ok := m2.Get(e1.key)
			if !ok {
				return false
			}
			v1 := e1.value
			if isNil(v1) {
				if !isNil(v2) {
					return false
				}
			} else if !v1.Equals(v2) {
				return false
			}
		}
	}
	return true
}

// HashCode returns the hash code for the Map object.
func (m *Map[K, V]) HashCode() int {
	h := 0
	for _, bucket := range m.data {
		for _, e := range bucket {
			var keyHash, valueHash int
			if !isNil(e.key) {
				keyHash = e.key.HashCode()
			}
			if !isNil(e.value) {
				valueHash = e.value.HashCode()
			}
			h += keyHash ^ valueHash
		}
	}
	return h
}
