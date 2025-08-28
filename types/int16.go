package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// Int16 is a wrapper for the built-in int16 type, implementing the Object interface.
type Int16 struct {
	object.BaseObject
	Value int16
}

// NewInt16 creates a new Int16 object.
func NewInt16(value int16) *Int16 {
	return &Int16{Value: value}
}

// String returns the string representation of the Int16 object.
func (i *Int16) String() string {
	return strconv.FormatInt(int64(i.Value), 10)
}

// Equals checks if two Int16 objects are equal.
func (i *Int16) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if i2, ok := other.(*Int16); ok {
		return i.Value == i2.Value
	}
	return false
}

// HashCode returns the hash code for the Int16 object.
func (i *Int16) HashCode() int {
	return int(i.Value)
}

// IntValue returns the value of the number as an int.
func (i *Int16) IntValue() int {
	return int(i.Value)
}

// Int64Value returns the value of the number as an int64.
func (i *Int16) Int64Value() int64 {
	return int64(i.Value)
}

// Float32Value returns the value of the number as a float32.
func (i *Int16) Float32Value() float32 {
	return float32(i.Value)
}

// Float64Value returns the value of the number as a float64.
func (i *Int16) Float64Value() float64 {
	return float64(i.Value)
}

// CompareTo compares this object with the specified object for order.
func (i *Int16) CompareTo(other *Int16) int {
	if i.Value < other.Value {
		return -1
	}
	if i.Value > other.Value {
		return 1
	}
	return 0
}
