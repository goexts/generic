// Package examples demonstrates set package operations.
package examples

import (
	"fmt"
	"sort"

	"github.com/goexts/generic/set"
)

func ExampleContains() {
	s := []int{1, 2, 3}
	fmt.Println("contains 2:", set.Contains(s, 2))
	fmt.Println("contains 4:", set.Contains(s, 4))

	// Output:
	// contains 2: true
	// contains 4: false
}

func ExampleUnique() {
	s := []int{1, 2, 2, 3, 3, 3}
	unique := set.Unique(s)
	sort.Ints(unique) // Sort the result for consistent output
	fmt.Println("unique elements:", unique)

	// Output:
	// unique elements: [1 2 3]
}
