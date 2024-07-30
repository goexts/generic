// Package types defines interfaces for different types of data.
package types

// Const is an interface that represents a value that is either a Number, a Boolean, or a string.
// It is used to define constants in Go.
type Const interface {
	Number | Boolean | ~string
}
