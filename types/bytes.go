package types

import (
	"bytes"

	"github.com/goexts/generic/object"
)

// Bytes is a wrapper for the built-in []byte type, implementing the Object interface.
type Bytes struct {
	object.BaseObject
	Value []byte
}

// NewBytes creates a new Bytes object.
func NewBytes(value []byte) *Bytes {
	return &Bytes{Value: value}
}

// String returns the string representation of the Bytes object.
func (b *Bytes) String() string {
	return string(b.Value)
}

// Equals checks if two Bytes objects are equal.
func (b *Bytes) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	if b2, ok := other.(*Bytes); ok {
		return bytes.Equal(b.Value, b2.Value)
	}
	return false
}

// HashCode returns the hash code for the Bytes object.
func (b *Bytes) HashCode() int {
	h := 1
	for _, v := range b.Value {
		h = 31*h + int(v)
	}
	return h
}
