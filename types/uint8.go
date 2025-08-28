package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// UInt8 is a wrapper for the built-in uint8 type, implementing the Object interface.
type UInt8 struct {
	object.BaseObject
	Value uint8
}

// NewUInt8 creates a new UInt8 object.
func NewUInt8(value uint8) *UInt8 {
	return &UInt8{Value: value}
}

// String returns the string representation of the UInt8 object.
func (u *UInt8) String() string {
	return strconv.FormatUint(uint64(u.Value), 10)
}

// Equals checks if two UInt8 objects are equal.
func (u *UInt8) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if u2, ok := other.(*UInt8); ok {
		return u.Value == u2.Value
	}
	return false
}

// HashCode returns the hash code for the UInt8 object.
func (u *UInt8) HashCode() int {
	return int(u.Value)
}

// IntValue returns the value of the number as an int.
func (u *UInt8) IntValue() int {
	return int(u.Value)
}

// Int64Value returns the value of the number as an int64.
func (u *UInt8) Int64Value() int64 {
	return int64(u.Value)
}

// Float32Value returns the value of the number as a float32.
func (u *UInt8) Float32Value() float32 {
	return float32(u.Value)
}

// Float64Value returns the value of the number as a float64.
func (u *UInt8) Float64Value() float64 {
	return float64(u.Value)
}

// CompareTo compares this object with the specified object for order.
func (u *UInt8) CompareTo(other *UInt8) int {
	if u.Value < other.Value {
		return -1
	}
	if u.Value > other.Value {
		return 1
	}
	return 0
}
