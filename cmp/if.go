// Package cmp implements the functions, types, and interfaces for the module.
package cmp

// If function takes a boolean condition and two values of any type T,
// and returns the first value if the condition is true,
// and the second value if the condition is false.
func If[T any](condition bool, trueVal T, falseVal T) T {
	// If the condition is true, return the first value
	if condition {
		return trueVal
	}
	// Otherwise, return the second value
	return falseVal
}

// IfFunc function takes a boolean condition,
// a function to execute if the condition is true,
// and a function to execute if the condition is false.
// It returns the result of the function that was executed.
func IfFunc[T any](condition bool, trueFn func() T, falseFn func() T) T {
	// If the condition is true, execute the trueFn function
	if condition {
		return trueFn()
	}
	// Return the result of the falseFn function
	return falseFn()
}

// IfFuncWithError function takes a boolean condition,
// a function to execute if the condition is true,
// and a function to execute if the condition is false.
// It returns the result of the function that was executed.
func IfFuncWithError[T any](condition bool, trueFn func() (T, error), falseFn func() (T, error)) (T, error) {
	// If the condition is true, execute the trueFn function and return its result
	if condition {
		return trueFn()
	}
	// If the condition is false, execute the falseFn function and return its result
	return falseFn()
}
