package types

import (
	"strconv"

	"github.com/goexts/generic/object"
)

// UIntPtr is a wrapper for the built-in uintptr type, implementing the Object interface.
type UIntPtr struct {
	object.BaseObject
	Value uintptr
}

// NewUIntPtr creates a new UIntPtr object.
func NewUIntPtr(value uintptr) *UIntPtr {
	return &UIntPtr{Value: value}
}

// String returns the string representation of the UIntPtr object.
func (u *UIntPtr) String() string {
	return "0x" + strconv.FormatUint(uint64(u.Value), 16)
}

// Equals checks if two UIntPtr objects are equal.
func (u *UIntPtr) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if u2, ok := other.(*UIntPtr); ok {
		return u.Value == u2.Value
	}
	return false
}

// HashCode returns the hash code for the UIntPtr object.
func (u *UIntPtr) HashCode() int {
	return int(u.Value)
}
