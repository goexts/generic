/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package generic implements the functions, types, and interfaces for the module.
package generic

// ChooseRight returns the rightmost value, ignoring the left value.
// This function is useful when you need to prioritize the right value over the left.
func ChooseRight[L any, R any](_ L, r R) R {
	return r
}

// ChooseLeft returns the leftmost value, ignoring the right value.
// This function is useful when you need to prioritize the left value over the right.
func ChooseLeft[L any, R any](l L, _ R) L {
	return l
}

// OnlyLeft returns the leftmost value, ignoring all other values.
// This function is useful when you need to extract the leftmost value from a tuple.
func OnlyLeft[L any, R1 any, R2 any](l L, _ R1, _ R2) L {
	return l
}

// OnlyRight returns the rightmost value, ignoring all other values.
// This function is useful when you need to extract the rightmost value from a tuple.
func OnlyRight[L1 any, L2 any, R any](_ L1, _ L2, r R) R {
	return r
}

// OnlyMiddle returns the middle value, ignoring all other values.
// This function is useful when you need to extract the middle value from a tuple.
func OnlyMiddle[L any, M any, R any](_ L, m M, _ R) M {
	return m
}
