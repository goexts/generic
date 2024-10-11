package types

// Zero is the zero value for a type.
func Zero[T any]() (zero T) {
	return zero
}

// ZeroOr returns def if v is the zero value.
func ZeroOr[T comparable](v T, def T) T {
	if v == Zero[T]() {
		return def
	}
	return v
}
