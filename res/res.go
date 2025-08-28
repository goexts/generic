package res

import "fmt"

// Result is a type that represents either a value of type T or an error.
// It is a monadic type, similar to Rust's Result or Haskell's Either.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a new successful Result with the given value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// Err creates a new failed Result with the given error.
func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{value: zero, err: err}
}

// Of converts a standard Go (value, error) pair into a Result.
func Of[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// IsOk returns true if the result is Ok.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the result is an error.
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Unwrap returns the value if the result is Ok, otherwise it panics.
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(fmt.Sprintf("called `Result.Unwrap()` on an `Err` value: %v", r.err))
	}
	return r.value
}

// UnwrapOr returns the contained Ok value or a provided default.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.value
}

// Expect returns the contained Ok value. Panics with a custom message if the value is an error.
func (r Result[T]) Expect(message string) T {
	if r.IsErr() {
		panic(fmt.Sprintf("%s: %v", message, r.err))
	}
	return r.value
}

// Ok returns the contained value and a boolean indicating if it was Ok.
// This is useful for safely accessing the value in a way that is idiomatic to Go.
func (r Result[T]) Ok() (T, bool) {
	return r.value, r.IsOk()
}

// Err returns the contained error, or nil if the result is Ok.
func (r Result[T]) Err() error {
	return r.err
}
