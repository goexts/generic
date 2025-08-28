/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package generic implements the functions, types, and interfaces for the module.
package generic

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
