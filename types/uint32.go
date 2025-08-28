package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// UInt32 is a wrapper for the built-in uint32 type, implementing the Object interface.
type UInt32 struct {
	object.BaseObject
	Value uint32
}

// NewUInt32 creates a new UInt32 object.
func NewUInt32(value uint32) *UInt32 {
	return &UInt32{Value: value}
}

// String returns the string representation of the UInt32 object.
func (u *UInt32) String() string {
	return strconv.FormatUint(uint64(u.Value), 10)
}

// Equals checks if two UInt32 objects are equal.
func (u *UInt32) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if u2, ok := other.(*UInt32); ok {
		return u.Value == u2.Value
	}
	return false
}

// HashCode returns the hash code for the UInt32 object.
func (u *UInt32) HashCode() int {
	return int(u.Value)
}

// IntValue returns the value of the number as an int.
func (u *UInt32) IntValue() int {
	return int(u.Value)
}

// Int64Value returns the value of the number as an int64.
func (u *UInt32) Int64Value() int64 {
	return int64(u.Value)
}

// Float32Value returns the value of the number as a float32.
func (u *UInt32) Float32Value() float32 {
	return float32(u.Value)
}

// Float64Value returns the value of the number as a float64.
func (u *UInt32) Float64Value() float64 {
	return float64(u.Value)
}

// CompareTo compares this object with the specified object for order.
func (u *UInt32) CompareTo(other *UInt32) int {
	if u.Value < other.Value {
		return -1
	}
	if u.Value > other.Value {
		return 1
	}
	return 0
}
