// Package configure provides utilities for applying functional options to objects.
package configure

// WithValidation creates an option that validates the target object.
// If the validator function returns an error, the configuration process will stop.
func WithValidation[T any](validator func(*T) error) OptionE[T] {
	return func(t *T) error {
		return validator(t)
	}
}
