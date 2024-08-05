// Package types defines interfaces for different types of data.
package types

// String is an interface that represents a string-like type.
// It can be a string, a byte slice, or a rune slice.
type String interface {
	~string | ~[]byte | ~[]rune
	String() string
}
