package must

import "fmt"

// Must panics if err is not nil, otherwise it returns the value v.
// It is a convenience wrapper for function calls that return (T, error)
// in contexts where the error is considered a fatal, non-recoverable bug.
// See the package documentation for appropriate use cases.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Must2 is the same as Must, but for functions that return (T, U, error).
func Must2[T any, U any](v1 T, v2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

// Cast performs a type assertion and panics if it fails.
// It provides a more informative panic message than a raw type assertion.
func Cast[T any](v any) T {
	if ret, ok := v.(T); ok {
		return ret
	}
	var zero T
	panic(fmt.Sprintf("type assertion failed: value of type %T cannot be cast to %T", v, zero))
}
