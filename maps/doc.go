/*
Package maps provides a set of generic functions for common operations on maps.

This package is a generated adapter and mirrors the public API of the standard
Go experimental package `golang.org/x/exp/maps`. It offers a convenient way
to access these common utilities, such as cloning maps, extracting keys or
values, and comparing maps for equality.

Example (Extracting Keys and Values):

	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Carol": 35,
	}

	names := maps.Keys(ages)   // Returns ["Alice", "Bob", "Carol"] (order is not guaranteed)
	_ = maps.Values(ages) // Returns [30, 25, 35] (order is not guaranteed)

	// Sort the names for deterministic output
	slices.Sort(names)
	// names is now ["Alice", "Bob", "Carol"]
*/
package maps
