package types

import (
	"math"
	"strconv"

	"github.com/goexts/generic/object"
)

// Complex128 is a wrapper for the built-in complex128 type, implementing the Object interface.
type Complex128 struct {
	object.BaseObject
	Value complex128
}

// NewComplex128 creates a new Complex128 object.
func NewComplex128(value complex128) *Complex128 {
	return &Complex128{Value: value}
}

// String returns the string representation of the Complex128 object.
func (c *Complex128) String() string {
	return strconv.FormatComplex(c.Value, 'f', -1, 128)
}

// Equals checks if two Complex128 objects are equal.
func (c *Complex128) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if c2, ok := other.(*Complex128); ok {
		return c.Value == c2.Value
	}
	return false
}

// HashCode returns the hash code for the Complex128 object.
func (c *Complex128) HashCode() int {
	realPart := math.Float64bits(real(c.Value))
	imagPart := math.Float64bits(imag(c.Value))
	// Inspired by Java's Double.hashCode()
	xor := realPart ^ imagPart
	return int(xor ^ (xor >> 32))
}
