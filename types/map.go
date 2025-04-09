// Package types implements the functions, types, and interfaces for the module.
package types

type Map[K comparable, V any] interface {
	Object
	~map[K]V
}
