package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// UInt is a wrapper for the built-in uint type, implementing the Object interface.
type UInt struct {
	object.BaseObject
	Value uint
}

// NewUInt creates a new UInt object.
func NewUInt(value uint) *UInt {
	return &UInt{Value: value}
}

// String returns the string representation of the UInt object.
func (u *UInt) String() string {
	return strconv.FormatUint(uint64(u.Value), 10)
}

// Equals checks if two UInt objects are equal.
func (u *UInt) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if u2, ok := other.(*UInt); ok {
		return u.Value == u2.Value
	}
	return false
}

// HashCode returns the hash code for the UInt object.
func (u *UInt) HashCode() int {
	return int(u.Value)
}

// IntValue returns the value of the number as an int.
func (u *UInt) IntValue() int {
	return int(u.Value)
}

// Int64Value returns the value of the number as an int64.
func (u *UInt) Int64Value() int64 {
	return int64(u.Value)
}

// Float32Value returns the value of the number as a float32.
func (u *UInt) Float32Value() float32 {
	return float32(u.Value)
}

// Float64Value returns the value of the number as a float64.
func (u *UInt) Float64Value() float64 {
	return float64(u.Value)
}

// CompareTo compares this object with the specified object for order.
func (u *UInt) CompareTo(other *UInt) int {
	if u.Value < other.Value {
		return -1
	}
	if u.Value > other.Value {
		return 1
	}
	return 0
}
