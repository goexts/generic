package must

import "fmt"

// Do panics if err is not nil, otherwise it returns the value v.
// It is useful for wrapping function calls that return a value and an error,
// where the error is not expected.
func Do[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Do2 is similar to Do, but for functions that return two values and an error.
func Do2[T any, U any](v1 T, v2 U, err error) (T, U) {
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
