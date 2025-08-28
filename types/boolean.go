package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// Boolean is a wrapper for the built-in bool type, implementing the Object interface.
type Boolean struct {
	object.BaseObject
	Value bool
}

// NewBoolean creates a new Boolean object.
func NewBoolean(value bool) *Boolean {
	return &Boolean{Value: value}
}

// String returns the string representation of the Boolean object.
func (b *Boolean) String() string {
	return strconv.FormatBool(b.Value)
}

// Equals checks if two Boolean objects are equal.
func (b *Boolean) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if b2, ok := other.(*Boolean); ok {
		return b.Value == b2.Value
	}
	return false
}

// HashCode returns the hash code for the Boolean object.
func (b *Boolean) HashCode() int {
	if b.Value {
		return 1
	}
	return 0
}
