package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// Int is a wrapper for the built-in int type, implementing the Object interface.
type Int struct {
	object.BaseObject
	Value int
}

// NewInt creates a new Int object.
func NewInt(value int) *Int {
	return &Int{Value: value}
}

// String returns the string representation of the Int object.
func (i *Int) String() string {
	return strconv.Itoa(i.Value)
}

// Equals checks if two Int objects are equal.
func (i *Int) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if i2, ok := other.(*Int); ok {
		return i.Value == i2.Value
	}
	return false
}

// HashCode returns the hash code for the Int object.
func (i *Int) HashCode() int {
	return i.Value // For simplicity, use the int value itself as hash code.
}

// IntValue returns the value of the number as an int.
func (i *Int) IntValue() int {
	return i.Value
}

// Int64Value returns the value of the number as an int64.
func (i *Int) Int64Value() int64 {
	return int64(i.Value)
}

// Float32Value returns the value of the number as a float32.
func (i *Int) Float32Value() float32 {
	return float32(i.Value)
}

// Float64Value returns the value of the number as a float64.
func (i *Int) Float64Value() float64 {
	return float64(i.Value)
}

// CompareTo compares this object with the specified object for order.
func (i *Int) CompareTo(other *Int) int {
	if i.Value < other.Value {
		return -1
	}
	if i.Value > other.Value {
		return 1
	}
	return 0
}
