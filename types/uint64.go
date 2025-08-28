package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// UInt64 is a wrapper for the built-in uint64 type, implementing the Object interface.
type UInt64 struct {
	object.BaseObject
	Value uint64
}

// NewUInt64 creates a new UInt64 object.
func NewUInt64(value uint64) *UInt64 {
	return &UInt64{Value: value}
}

// String returns the string representation of the UInt64 object.
func (u *UInt64) String() string {
	return strconv.FormatUint(u.Value, 10)
}

// Equals checks if two UInt64 objects are equal.
func (u *UInt64) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if u2, ok := other.(*UInt64); ok {
		return u.Value == u2.Value
	}
	return false
}

// HashCode returns the hash code for the UInt64 object.
func (u *UInt64) HashCode() int {
	return int(u.Value ^ (u.Value >> 32))
}

// IntValue returns the value of the number as an int.
func (u *UInt64) IntValue() int {
	return int(u.Value)
}

// Int64Value returns the value of the number as an int64.
func (u *UInt64) Int64Value() int64 {
	return int64(u.Value)
}

// Float32Value returns the value of the number as a float32.
func (u *UInt64) Float32Value() float32 {
	return float32(u.Value)
}

// Float64Value returns the value of the number as a float64.
func (u *UInt64) Float64Value() float64 {
	return float64(u.Value)
}

// CompareTo compares this object with the specified object for order.
func (u *UInt64) CompareTo(other *UInt64) int {
	if u.Value < other.Value {
		return -1
	}
	if u.Value > other.Value {
		return 1
	}
	return 0
}
