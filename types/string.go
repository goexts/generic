package types

import (
	"github.com/goexts/generic/object"
)

// String is a wrapper for the built-in string type, implementing the Object interface.
type String struct {
	object.BaseObject
	Value string
}

// NewString creates a new String object.
func NewString(value string) *String {
	return &String{Value: value}
}

// String returns the string representation of the String object.
func (s *String) String() string {
	return s.Value
}

// Equals checks if two String objects are equal.
func (s *String) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if s2, ok := other.(*String); ok {
		return s.Value == s2.Value
	}
	return false
}

// HashCode returns the hash code for the String object.
// This implementation is inspired by Java's String.hashCode().
func (s *String) HashCode() int {
	h := 0
	for _, r := range s.Value {
		h = 31*h + int(r)
	}
	return h
}
