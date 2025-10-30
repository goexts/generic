/*
Package maps provides a set of generic functions for common operations on maps.

This package is a generated adapter that mirrors the public API of the standard
Go experimental package `golang.org/x/exp/maps`. It offers a convenient way
to access common utilities for cloning, extracting keys/values, and comparing maps.

# Usage

## Extracting Keys and Values

A common task is to get a slice of all keys or values from a map.

	userPermissions := map[string]string{"admin": "all", "editor": "write", "viewer": "read"}

	// Get all keys.
	keys := maps.Keys(userPermissions) // Result: []string{"admin", "editor", "viewer"} (order not guaranteed)

	// Get all values.
	vals := maps.Values(userPermissions) // Result: []string{"all", "write", "read"} (order not guaranteed)

	// For deterministic output, you can sort the resulting slices:
	// slices.Sort(keys)
	// slices.Sort(vals)

## Cloning a Map

Create a shallow copy of a map.

	original := map[string]int{"a": 1, "b": 2}
	clone := maps.Clone(original)
	clone["c"] = 3 // Modifying the clone does not affect the original.

## Comparing Maps

Check if two maps are equal (contain the same keys and values).

	map1 := map[string]int{"a": 1, "b": 2}
	map2 := map[string]int{"b": 2, "a": 1}
	map3 := map[string]int{"a": 1, "c": 3}

	equal12 := maps.Equal(map1, map2) // true
	equal13 := maps.Equal(map1, map3) // false

For more functions and detailed behavior, please refer to the documentation for
the standard `maps` package.
*/
package maps
