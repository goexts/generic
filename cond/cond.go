package cond

// If returns trueVal if the condition is true, and falseVal otherwise.
// This is a generic, eagerly-evaluated ternary-like expression. Both trueVal
// and falseVal are evaluated before this function is called.
//
// For lazy evaluation where values are expensive to compute, see IfFunc.
func If[T any](condition bool, trueVal T, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// IfFunc returns the result of trueFn if the condition is true, and the result
// of falseFn otherwise.
// This is a generic, lazily-evaluated ternary-like expression. Only the function
// corresponding to the condition's outcome is executed.
func IfFunc[T any](condition bool, trueFn func() T, falseFn func() T) T {
	if condition {
		return trueFn()
	}
	return falseFn()
}

// IfFuncE returns the result of trueFn if the condition is true, and the result
// of falseFn otherwise. It is the error-returning version of IfFunc.
// This is a generic, lazily-evaluated ternary-like expression. Only the function
// corresponding to the condition's outcome is executed.
func IfFuncE[T any](condition bool, trueFn func() (T, error), falseFn func() (T, error)) (T, error) {
	if condition {
		return trueFn()
	}
	return falseFn()
}
