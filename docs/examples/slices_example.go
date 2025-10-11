// Package main provides examples for using the slices package.
// These examples demonstrate common operations like mapping, filtering, and reducing slices.
package main

import (
	"fmt"

	"github.com/goexts/generic/slices"
)

func main() {
	// Example 1: Basic operations
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Map: Double each number
	doubled := slices.Map(nums, func(x int) int {
		return x * 2
	})
	fmt.Println("Doubled:", doubled) // [2 4 6 8 10 12 14 16 18 20]

	// Filter: Get even numbers
	evens := slices.Filter(doubled, func(x int) bool {
		return x%2 == 0
	})
	fmt.Println("Even numbers:", evens) // [2 4 6 8 10 12 14 16 18 20]

	// Reduce: Calculate sum
	sum := slices.Reduce(nums, 0, func(acc, x int) int {
		return acc + x
	})
	fmt.Println("Sum:", sum) // 55

	// Example 2: Working with strings
	names := []string{"alice", "bob", "charlie", "dave"}

	// Map: Capitalize names
	capitalized := slices.Map(names, func(s string) string {
		if len(s) == 0 {
			return s
		}
		return string(append([]byte{s[0] - 32}, s[1:]...))
	})
	fmt.Println("Capitalized:", capitalized) // [Alice Bob Charlie Dave]

	// Example 3: Using FilterIncluded and FilterExcluded
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evensOnly := slices.FilterIncluded(allNumbers, []int{2, 4, 6, 8, 10})
	fmt.Println("Even numbers (FilterIncluded):", evensOnly) // [2 4 6 8 10]

	oddsOnly := slices.FilterExcluded(allNumbers, []int{2, 4, 6, 8, 10})
	fmt.Println("Odd numbers (FilterExcluded):", oddsOnly) // [1 3 5 7 9]
}
