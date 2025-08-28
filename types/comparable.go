package types

import "github.com/goexts/generic/object"

// Comparable is an interface for objects that can be compared to another object of the same type.
// It is the equivalent of java.lang.Comparable.
type Comparable[T object.Object] interface {
	object.Object

	// CompareTo compares this object with the specified object for order.
	// Returns a negative integer, zero, or a positive integer as this object
	// is less than, equal to, or greater than the specified object.
	CompareTo(other T) int
}
