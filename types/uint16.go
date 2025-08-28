package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// UInt16 is a wrapper for the built-in uint16 type, implementing the Object interface.
type UInt16 struct {
	object.BaseObject
	Value uint16
}

// NewUInt16 creates a new UInt16 object.
func NewUInt16(value uint16) *UInt16 {
	return &UInt16{Value: value}
}

// String returns the string representation of the UInt16 object.
func (u *UInt16) String() string {
	return strconv.FormatUint(uint64(u.Value), 10)
}

// Equals checks if two UInt16 objects are equal.
func (u *UInt16) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if u2, ok := other.(*UInt16); ok {
		return u.Value == u2.Value
	}
	return false
}

// HashCode returns the hash code for the UInt16 object.
func (u *UInt16) HashCode() int {
	return int(u.Value)
}

// IntValue returns the value of the number as an int.
func (u *UInt16) IntValue() int {
	return int(u.Value)
}

// Int64Value returns the value of the number as an int64.
func (u *UInt16) Int64Value() int64 {
	return int64(u.Value)
}

// Float32Value returns the value of the number as a float32.
func (u *UInt16) Float32Value() float32 {
	return float32(u.Value)
}

// Float64Value returns the value of the number as a float64.
func (u *UInt16) Float64Value() float64 {
	return float64(u.Value)
}

// CompareTo compares this object with the specified object for order.
func (u *UInt16) CompareTo(other *UInt16) int {
	if u.Value < other.Value {
		return -1
	}
	if u.Value > other.Value {
		return 1
	}
	return 0
}
