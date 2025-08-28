// Package cond provides utility functions for conditional logic, like ternary operators.
package cond

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
//
// Decrypted: use IfFuncE instead of IfFuncWithError
func IfFuncWithError[T any](condition bool, trueFn func() (T, error), falseFn func() (T, error)) (T, error) {
	// If the condition is true, execute the trueFn function and return its result
	if condition {
		return trueFn()
	}
	// If the condition is false, execute the falseFn function and return its result
	return falseFn()
}

func IfFuncE[T any](condition bool, trueFn func() (T, error), falseFn func() (T, error)) (T, error) {
	// If the condition is true, execute the trueFn function and return its result
	if condition {
		return trueFn()
	}
	// If the condition is false, execute the falseFn function and return its result
	return falseFn()
}

// Choose returns the left value if cond is true, otherwise it returns the right value.
// This function is useful when you need to conditionally choose between two values.
func Choose[T any](cond bool, l, r T) T {
	if cond {
		return l
	}
	return r
}

// ChooseFunc returns the result of the left function if cond is true, otherwise it returns the result of the right function.
// This function is useful when you need to conditionally choose between two functions.
func ChooseFunc[T any](cond bool, l, r func() T) T {
	if cond {
		return l()
	}
	return r()
}

// ChooseRight returns the rightmost value, ignoring the left value.
// This function is useful when you need to prioritize the right value over the left.
func ChooseRight[L, R any](_ L, r R) R {
	return r
}

// ChooseLeft returns the leftmost value, ignoring the right value.
// This function is useful when you need to prioritize the left value over the right.
func ChooseLeft[L, R any](l L, _ R) L {
	return l
}

// OnlyLeft returns the leftmost value, ignoring all other values.
// This function is useful when you need to extract the leftmost value from a tuple.
func OnlyLeft[L, R1, R2 any](l L, _ R1, _ R2) L {
	return l
}

// OnlyRight returns the rightmost value, ignoring all other values.
// This function is useful when you need to extract the rightmost value from a tuple.
func OnlyRight[L1, L2, R any](_ L1, _ L2, r R) R {
	return r
}

// OnlyMiddle returns the middle value, ignoring all other values.
// This function is useful when you need to extract the middle value from a tuple.
func OnlyMiddle[L, M, R any](_ L, m M, _ R) M {
	return m
}
