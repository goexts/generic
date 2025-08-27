/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package generic implements the functions, types, and interfaces for the module.
package generic

// Difference returns the unique elements from slice2 that are not present in slice1.
// It effectively computes the set difference `slice2 - slice1`.
// The order of elements in the returned slice is not guaranteed.
func Difference[T comparable](slice1 []T, slice2 []T) []T {
	// Create a set of elements to be excluded (from slice1) for efficient lookup.
	excludeSet := make(map[T]struct{}, len(slice1))
	for _, item := range slice1 {
		excludeSet[item] = struct{}{}
	}

	// Use a map to store the result to handle duplicates in slice2 and ensure uniqueness.
	resultSet := make(map[T]struct{})
	for _, item := range slice2 {
		// If the item from slice2 is NOT in the exclude set, add it to the result.
		if _, found := excludeSet[item]; !found {
			resultSet[item] = struct{}{}
		}
	}

	// Convert the result set map back to a slice.
	result := make([]T, 0, len(resultSet))
	for item := range resultSet {
		result = append(result, item)
	}
	return result
}
