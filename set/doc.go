/*
Package set provides a collection of generic, stateless functions for performing
set-like operations on standard Go slices.

# Core Philosophy

Instead of introducing a new, stateful `Set` data structure, this package provides
utilities that treat slices as if they were sets. This approach offers several
advantages:

  - **Lightweight**: It requires no new types, reducing the learning curve.
  - **Idiomatic**: It integrates seamlessly with the rest of the Go ecosystem, which
    heavily relies on slices.
  - **Functional**: It promotes a functional style of programming by taking slices
    as input and returning new slices as output.

The functions in this package operate on slices of any `comparable` type.

# Usage

Note: The output of these functions is not guaranteed to be in any specific order.
If a stable order is required, sort the resulting slice.

	setA := []int{1, 2, 3, 4}
	setB := []int{3, 4, 5, 6}

## Union

`Union` returns a new slice containing all unique elements from the input slices.

	union := set.Union(setA, setB) // Result (order not guaranteed): [1, 2, 3, 4, 5, 6]

## Intersection

`Intersection` returns a new slice containing only the elements that exist in
all input slices.

	intersection := set.Intersection(setA, setB) // Result (order not guaranteed): [3, 4]

## Difference

`Difference` returns a new slice containing elements from the first slice that
are not present in the second slice.

	diffA_B := set.Difference(setA, setB) // Result: [1, 2]
	diffB_A := set.Difference(setB, setA) // Result: [5, 6]

## Unique

`Unique` removes duplicate elements from a single slice.

	items := []string{"a", "b", "a", "c", "b"}
	unique := set.Unique(items) // Result (order not guaranteed): ["a", "b", "c"]
*/
package set
