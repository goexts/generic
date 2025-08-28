package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// Int64 is a wrapper for the built-in int64 type, implementing the Object interface.
type Int64 struct {
	object.BaseObject
	Value int64
}

// NewInt64 creates a new Int64 object.
func NewInt64(value int64) *Int64 {
	return &Int64{Value: value}
}

// String returns the string representation of the Int64 object.
func (i *Int64) String() string {
	return strconv.FormatInt(i.Value, 10)
}

// Equals checks if two Int64 objects are equal.
func (i *Int64) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if i2, ok := other.(*Int64); ok {
		return i.Value == i2.Value
	}
	return false
}

// HashCode returns the hash code for the Int64 object.
func (i *Int64) HashCode() int {
	// Inspired by Java's Long.hashCode()
	return int(i.Value ^ (i.Value >> 32))
}

// IntValue returns the value of the number as an int.
func (i *Int64) IntValue() int {
	return int(i.Value)
}

// Int64Value returns the value of the number as an int64.
func (i *Int64) Int64Value() int64 {
	return i.Value
}

// Float32Value returns the value of the number as a float32.
func (i *Int64) Float32Value() float32 {
	return float32(i.Value)
}

// Float64Value returns the value of the number as a float64.
func (i *Int64) Float64Value() float64 {
	return float64(i.Value)
}

// CompareTo compares this object with the specified object for order.
func (i *Int64) CompareTo(other *Int64) int {
	if i.Value < other.Value {
		return -1
	}
	if i.Value > other.Value {
		return 1
	}
	return 0
}
