package types

import (
	"math"
	"strconv"

	"github.com/goexts/generic/object"
)

// Float32 is a wrapper for the built-in float32 type, implementing the Object interface.
type Float32 struct {
	object.BaseObject
	Value float32
}

// NewFloat32 creates a new Float32 object.
func NewFloat32(value float32) *Float32 {
	return &Float32{Value: value}
}

// String returns the string representation of the Float32 object.
func (f *Float32) String() string {
	return strconv.FormatFloat(float64(f.Value), 'f', -1, 32)
}

// Equals checks if two Float32 objects are equal.
func (f *Float32) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if f2, ok := other.(*Float32); ok {
		return f.Value == f2.Value
	}
	return false
}

// HashCode returns the hash code for the Float32 object.
func (f *Float32) HashCode() int {
	// Use math.Float32bits to get a stable integer representation of the float's bits.
	return int(math.Float32bits(f.Value))
}

// IntValue returns the value of the number as an int.
func (f *Float32) IntValue() int {
	return int(f.Value)
}

// Int64Value returns the value of the number as an int64.
func (f *Float32) Int64Value() int64 {
	return int64(f.Value)
}

// Float32Value returns the value of the number as a float32.
func (f *Float32) Float32Value() float32 {
	return f.Value
}

// Float64Value returns the value of the number as a float64.
func (f *Float32) Float64Value() float64 {
	return float64(f.Value)
}

// CompareTo compares this object with the specified object for order.
func (f *Float32) CompareTo(other *Float32) int {
	if f.Value < other.Value {
		return -1
	}
	if f.Value > other.Value {
		return 1
	}
	return 0
}
