package types

import (
	"math"
	"strconv"

	"github.com/goexts/generic/object"
)

// Float64 is a wrapper for the built-in float64 type, implementing the Object interface.
type Float64 struct {
	object.BaseObject
	Value float64
}

// NewFloat64 creates a new Float64 object.
func NewFloat64(value float64) *Float64 {
	return &Float64{Value: value}
}

// String returns the string representation of the Float64 object.
func (f *Float64) String() string {
	return strconv.FormatFloat(f.Value, 'f', -1, 64)
}

// Equals checks if two Float64 objects are equal.
func (f *Float64) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if f2, ok := other.(*Float64); ok {
		return f.Value == f2.Value
	}
	return false
}

// HashCode returns the hash code for the Float64 object.
func (f *Float64) HashCode() int {
	// Use math.Float64bits to get a stable integer representation of the float's bits.
	return int(math.Float64bits(f.Value))
}

// IntValue returns the value of the number as an int.
func (f *Float64) IntValue() int {
	return int(f.Value)
}

// Int64Value returns the value of the number as an int64.
func (f *Float64) Int64Value() int64 {
	return int64(f.Value)
}

// Float32Value returns the value of the number as a float32.
func (f *Float64) Float32Value() float32 {
	return float32(f.Value)
}

// Float64Value returns the value of the number as a float64.
func (f *Float64) Float64Value() float64 {
	return f.Value
}

// CompareTo compares this object with the specified object for order.
func (f *Float64) CompareTo(other *Float64) int {
	if f.Value < other.Value {
		return -1
	}
	if f.Value > other.Value {
		return 1
	}
	return 0
}
