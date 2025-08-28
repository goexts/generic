package ptr

// Of returns a pointer to the given value v.
// This is a convenient helper for creating a pointer to a literal or a variable
// in a single expression, often used for struct field initialization.
//
// Example:
//
//	config := &MyConfig{ Timeout: ptr.Of(5*time.Second) }
func Of[T any](v T) *T {
	return &v
}

// Val returns the value that the pointer v points to.
// If the pointer is nil, it safely returns the zero value of the type T.
// This avoids a panic when dereferencing a nil pointer.
func Val[T any](v *T) T {
	if v != nil {
		return *v
	}
	var zero T
	return zero
}

// To converts an `any` value to a pointer of type *T.
// It handles three cases:
//  1. If v is of type T, it returns a pointer to it (&v).
//  2. If v is already of type *T, it returns v directly.
//  3. Otherwise, it returns a new pointer to a zero value of T.
func To[T any](v any) *T {
	// Check if v is of type T and return its address
	if val, ok := v.(T); ok {
		return &val
	}
	// Check if v is already a *T pointer and return it
	if ptr, ok := v.(*T); ok {
		return ptr
	}
	return new(T)
}

// ToVal converts an `any` value to a value of type T.
// It handles three cases:
//  1. If v is a non-nil pointer of type *T, it returns the dereferenced value.
//  2. If v is of type T, it returns v directly.
//  3. Otherwise, it returns the zero value of T.
func ToVal[T any](v any) T {
	// Handle non-nil pointer case
	if ptr, ok := v.(*T); ok && ptr != nil {
		return *ptr
	}
	// Handle direct value case
	if val, ok := v.(T); ok {
		return val
	}
	// Return zero value as fallback
	var zero T
	return zero
}
