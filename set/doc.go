/*
Package set provides a collection of generic, stateless functions for performing
set-like operations on standard Go slices.

# Core Philosophy

Instead of introducing a new `Set` data structure, this package provides
utilities that treat slices as sets. This approach offers several advantages:

  - It is lightweight and requires no new types to learn.
  - It integrates seamlessly with the rest of the Go ecosystem, which heavily
    relies on slices.
  - It promotes a functional style of programming.

The functions in this package operate on slices of any `comparable` type.

# Example

	setA := []int{1, 2, 3, 4}
	setB := []int{3, 4, 5, 6}

	// Calculate the union of the two sets
	union := set.Union(setA, setB) // Result (order not guaranteed): [1, 2, 3, 4, 5, 6]

	// Calculate the intersection of the two sets
	intersection := set.Intersection(setA, setB) // Result (order not guaranteed): [3, 4]

	// Calculate the difference (elements in B but not in A)
	difference := set.Difference(setA, setB) // Result (order not guaranteed): [5, 6]

	// Remove duplicates from a slice
	unique := set.Unique([]string{"a", "b", "a", "c", "b"}) // Result (order not guaranteed): ["a", "b", "c"]
*/
package set
