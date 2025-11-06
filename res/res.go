package res

import "fmt"

// Result is a type that represents either a success (containing a value of type T)
// or a failure (containing an error). It is a monadic type that allows for
// chaining operations in a clean, readable way.
// See the package documentation for more details and usage examples.
type Result[T any] struct {
	value T
	err   error
}

// Ok creates a new successful Result containing the given value.
func Ok[T any](value T) Result[T] {
	return Result[T]{value: value}
}

// Err creates a new failed Result containing the given error.
func Err[T any](err error) Result[T] {
	var zero T
	return Result[T]{value: zero, err: err}
}

// Of converts a standard Go (value, error) pair into a Result.
// If err is not nil, it returns an Err result; otherwise, it returns an Ok result.
func Of[T any](value T, err error) Result[T] {
	if err != nil {
		return Err[T](err)
	}
	return Ok(value)
}

// IsOk returns true if the result is Ok (i.e., does not contain an error).
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// IsErr returns true if the result is an Err (i.e., contains an error).
func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Unwrap returns the contained Ok value. It panics if the result is an Err.
// Because this function may panic, it should only be used when the caller is
// certain that the result is Ok, or when a panic is the desired behavior.
// See also: Expect, UnwrapOr.
func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic(fmt.Sprintf("called `Result.Unwrap()` on an `Err` value: %v", r.err))
	}
	return r.value
}

// UnwrapOr returns the contained Ok value or a provided default value.
// It is a safe way to access the value without panicking.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.IsErr() {
		return defaultValue
	}
	return r.value
}

// Unpack unpacks the Result into the original value and error pair.
//
// Return:
// - T: Included value (if Result is Ok)
// - error: Contains an error (if Result is Err)
//
// This method provides a way to convert the Result type back to a standard Go (value, error) pair,
// Facilitate interaction with existing code or APIs that expect this form of return.
func (r Result[T]) Unpack() (T, error) {
	return r.value, r.err
}

// Expect returns the contained Ok value. It panics with a custom message if
// the result is an Err.
// This is similar to Unwrap but provides a more context-specific panic message.
func (r Result[T]) Expect(message string) T {
	if r.IsErr() {
		panic(fmt.Sprintf("%s: %v", message, r.err))
	}
	return r.value
}

// Ok returns the contained value and a boolean indicating if the result was Ok.
// This provides a safe, idiomatic Go way to access the value.
func (r Result[T]) Ok() (T, bool) {
	return r.value, r.IsOk()
}

// Err returns the contained error, or nil if the result is Ok.
func (r Result[T]) Err() error {
	return r.err
}

// Or is a utility function that simplifies handling of (value, error) returns.
// It returns the value if err is nil, otherwise it returns the provided default value.
func Or[T any](v T, err error, defaultValue T) T {
	if err != nil {
		return defaultValue
	}
	return v
}

// OrZero is a utility function that simplifies handling of (value, error) returns.
// It returns the value if err is nil, otherwise it returns the zero value of the type.
func OrZero[T any](v T, err error) T {
	if err != nil {
		var zero T
		return zero
	}
	return v
}
