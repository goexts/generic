package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// Int8 is a wrapper for the built-in int8 type, implementing the Object interface.
type Int8 struct {
	object.BaseObject
	Value int8
}

// NewInt8 creates a new Int8 object.
func NewInt8(value int8) *Int8 {
	return &Int8{Value: value}
}

// String returns the string representation of the Int8 object.
func (i *Int8) String() string {
	return strconv.FormatInt(int64(i.Value), 10)
}

// Equals checks if two Int8 objects are equal.
func (i *Int8) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if i2, ok := other.(*Int8); ok {
		return i.Value == i2.Value
	}
	return false
}

// HashCode returns the hash code for the Int8 object.
func (i *Int8) HashCode() int {
	return int(i.Value)
}

// IntValue returns the value of the number as an int.
func (i *Int8) IntValue() int {
	return int(i.Value)
}

// Int64Value returns the value of the number as an int64.
func (i *Int8) Int64Value() int64 {
	return int64(i.Value)
}

// Float32Value returns the value of the number as a float32.
func (i *Int8) Float32Value() float32 {
	return float32(i.Value)
}

// Float64Value returns the value of the number as a float64.
func (i *Int8) Float64Value() float64 {
	return float64(i.Value)
}

// CompareTo compares this object with the specified object for order.
func (i *Int8) CompareTo(other *Int8) int {
	if i.Value < other.Value {
		return -1
	}
	if i.Value > other.Value {
		return 1
	}
	return 0
}
