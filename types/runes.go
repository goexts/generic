package types

import (
	"github.com/goexts/generic/object"
)

// Runes is a wrapper for the built-in []rune type, implementing the Object interface.
type Runes struct {
	object.BaseObject
	Value []rune
}

// NewRunes creates a new Runes object.
func NewRunes(value []rune) *Runes {
	return &Runes{Value: value}
}

// String returns the string representation of the Runes object.
func (r *Runes) String() string {
	return string(r.Value)
}

// Equals checks if two Runes objects are equal.
func (r *Runes) Equals(other object.Object) bool {
	if other == nil {
		return false
	}
	r2, ok := other.(*Runes)
	if !ok {
		return false
	}
	if len(r.Value) != len(r2.Value) {
		return false
	}
	for i, v := range r.Value {
		if v != r2.Value[i] {
			return false
		}
	}
	return true
}

// HashCode returns the hash code for the Runes object.
func (r *Runes) HashCode() int {
	h := 1
	for _, v := range r.Value {
		h = 31*h + int(v)
	}
	return h
}
