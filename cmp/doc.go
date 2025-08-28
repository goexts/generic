/*
Package cmp provides generic, type-safe functions for comparing ordered types.

This package offers fundamental comparison utilities like Min, Max, and Clamp.
Its central function, Compare, is designed to be fully compatible with the
`slices.SortFunc`, making it easy to sort slices of any ordered type.

Example (Sorting a custom type):

	type Product struct {
		ID    int
		Price float64
	}

	products := []Product{
		{ID: 1, Price: 99.99},
		{ID: 2, Price: 49.99},
		{ID: 3, Price: 74.99},
	}

	// Sort products by price using cmp.Compare
	slices.SortFunc(products, func(a, b Product) int {
		return cmp.Compare(a.Price, b.Price)
	})

	// products is now sorted by Price: [{2 49.99} {3 74.99} {1 99.99}]
*/
package cmp
