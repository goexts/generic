// Package set provides utility functions for treating slices as sets.
package set

// Contains checks whether a value is in a slice.
func Contains[T comparable](s []T, e T) bool {
	for _, t := range s {
		if t == e {
			return true
		}
	}
	return false
}

// Exists checks whether at least one element in a slice satisfies the predicate f.
func Exists[T any](s []T, f func(T) bool) bool {
	for _, t := range s {
		if f(t) {
			return true
		}
	}
	return false
}

// Intersection returns a new slice containing the elements that exist in both s1 and s2.
func Intersection[T comparable](s1 []T, s2 []T) []T {
	set := make(map[T]struct{}, len(s1))
	for _, t := range s1 {
		set[t] = struct{}{}
	}

	var result []T
	for _, t := range s2 {
		if _, ok := set[t]; ok {
			result = append(result, t)
		}
	}

	return result
}

// Difference returns a new slice containing the unique elements from s2 that are not present in s1.
// It effectively computes the set difference `s2 - s1`.
func Difference[T comparable](s1 []T, s2 []T) []T {
	// Create a set of elements to be excluded (from s1) for efficient lookup.
	excludeSet := make(map[T]struct{}, len(s1))
	for _, item := range s1 {
		excludeSet[item] = struct{}{}
	}

	// Use a map to store the result to handle duplicates in s2 and ensure uniqueness.
	resultSet := make(map[T]struct{})
	for _, item := range s2 {
		// If the item from s2 is NOT in the exclude set, add it to the result.
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
