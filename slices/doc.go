/*
Package slices provides a rich set of generic functions for common operations on
slices of any element type.

This package is a generated adapter and mirrors the public API of the standard
Go experimental package `golang.org/x/exp/slices`. It offers a convenient way
to access these common utilities for searching, sorting, comparing, and
manipulating slices.

For detailed information on the behavior of specific functions, please refer to
the official Go documentation for the `slices` package.

Example (Sorting and Searching):

	numbers := []int{3, 1, 4, 1, 5, 9}

	// Check if a slice contains a value
	_ = slices.Contains(numbers, 5) // true

	// Sort the slice in place
	slices.Sort(numbers)
	// numbers is now [1, 1, 3, 4, 5, 9]

	// Find the index of a value in a sorted slice
	idx, found := slices.BinarySearch(numbers, 4)
	// idx is 3, found is true
*/
package slices
