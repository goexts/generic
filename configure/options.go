// Package configure provides utilities for applying functional options to objects.
package configure

// Option represents a function that configures an object of type T.
// This is the core type for the Functional Options Pattern.
type Option[T any] func(*T)

// Apply allows Option[T] to satisfy the Applier[T] interface.
func (o Option[T]) Apply(target *T) {
	if o != nil {
		o(target)
	}
}

// OptionE represents a function that configures an object of type T and may return an error.
// 'E' stands for "Error".
type OptionE[T any] func(*T) error

// Apply allows OptionE[T] to satisfy the ApplierE[T] interface.
func (o OptionE[T]) Apply(target *T) error {
	if o != nil {
		return o(target)
	}
	return nil
}

// OptionConstraint is a generic interface that constrains a type to have an
// underlying type of func(*T). This is used to create flexible Apply functions.
type OptionConstraint[T any] interface {
	~func(*T)
}

// OptionEConstraint is a generic interface that constrains a type to have an
// underlying type of func(*T) error.
type OptionEConstraint[T any] interface {
	~func(*T) error
}
