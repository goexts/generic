/*
Package cmp provides generic, type-safe functions for comparing ordered types.

This package offers fundamental comparison utilities like Min, Max, and Clamp.
Its central function, Compare, is designed to be fully compatible with the standard
library's `slices.SortFunc`, making it easy to sort slices of any ordered type.

# Usage

Sorting a slice of structs by a specific field:

	type Product struct {
		ID    int
		Price float64
	}

	products := []Product{
		{ID: 1, Price: 99.99},
		{ID: 2, Price: 49.99},
		{ID: 3, Price: 74.99},
	}

	// Sort products by price in ascending order.
	slices.SortFunc(products, func(a, b Product) int {
		return cmp.Compare(a.Price, b.Price)
	})

Multi-level sorting with custom logic:

	type Employee struct {
		Department string
		Seniority  int
		Name       string
	}

	// Sort employees by Department (A-Z), then Seniority (descending), then Name (A-Z).
	func SortEmployees(employees []Employee) {
		slices.SortFunc(employees, func(a, b Employee) int {
			if res := cmp.Compare(a.Department, b.Department); res != 0 {
				return res
			}
			if res := cmp.Compare(b.Seniority, a.Seniority); res != 0 { // Note b, a for descending
				return res
			}
			return cmp.Compare(a.Name, b.Name)
		})
	}

For more details, refer to the function documentation.
*/
package cmp
