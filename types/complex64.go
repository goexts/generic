package types

import (
	"math"
	"strconv"

	"github.com/goexts/generic/object"
)

// Complex64 is a wrapper for the built-in complex64 type, implementing the Object interface.
type Complex64 struct {
	object.BaseObject
	Value complex64
}

// NewComplex64 creates a new Complex64 object.
func NewComplex64(value complex64) *Complex64 {
	return &Complex64{Value: value}
}

// String returns the string representation of the Complex64 object.
func (c *Complex64) String() string {
	return strconv.FormatComplex(complex128(c.Value), 'f', -1, 64)
}

// Equals checks if two Complex64 objects are equal.
func (c *Complex64) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if c2, ok := other.(*Complex64); ok {
		return c.Value == c2.Value
	}
	return false
}

// HashCode returns the hash code for the Complex64 object.
func (c *Complex64) HashCode() int {
	realPart := math.Float32bits(real(c.Value))
	imagPart := math.Float32bits(imag(c.Value))
	return int(realPart ^ imagPart)
}
