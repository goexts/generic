package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// Int32 is a wrapper for the built-in int32 type, implementing the Object interface.
type Int32 struct {
	object.BaseObject
	Value int32
}

// NewInt32 creates a new Int32 object.
func NewInt32(value int32) *Int32 {
	return &Int32{Value: value}
}

// String returns the string representation of the Int32 object.
func (i *Int32) String() string {
	return strconv.FormatInt(int64(i.Value), 10)
}

// Equals checks if two Int32 objects are equal.
func (i *Int32) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if i2, ok := other.(*Int32); ok {
		return i.Value == i2.Value
	}
	return false
}

// HashCode returns the hash code for the Int32 object.
func (i *Int32) HashCode() int {
	return int(i.Value)
}

// IntValue returns the value of the number as an int.
func (i *Int32) IntValue() int {
	return int(i.Value)
}

// Int64Value returns the value of the number as an int64.
func (i *Int32) Int64Value() int64 {
	return int64(i.Value)
}

// Float32Value returns the value of the number as a float32.
func (i *Int32) Float32Value() float32 {
	return float32(i.Value)
}

// Float64Value returns the value of the number as a float64.
func (i *Int32) Float64Value() float64 {
	return float64(i.Value)
}

// CompareTo compares this object with the specified object for order.
func (i *Int32) CompareTo(other *Int32) int {
	if i.Value < other.Value {
		return -1
	}
	if i.Value > other.Value {
		return 1
	}
	return 0
}
