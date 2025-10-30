// Package examples shows multi-level sort using cmp with slices.SortFunc (Report Q9).
package examples

import (
	"fmt"

	"github.com/goexts/generic/cmp"
	"github.com/goexts/generic/slices"
)

type Employee struct {
	Department string
	Seniority  int
	Name       string
}

// MultiLevelSort demonstrates sorting with multiple criteria
func MultiLevelSort(es []Employee) {
	slices.SortFunc(es, func(a, b Employee) int {
		if res := cmp.Compare(a.Department, b.Department); res != 0 {
			return res
		}
		if res := cmp.Compare(b.Seniority, a.Seniority); res != 0 {
			return res
		} // desc
		return cmp.Compare(a.Name, b.Name)
	})
}

func ExampleMultiLevelSort() {
	es := []Employee{
		{"Engineering", 3, "Alice"},
		{"Engineering", 5, "Bob"},
		{"HR", 1, "Carol"},
		{"Engineering", 5, "Aaron"},
	}

	MultiLevelSort(es)

	for _, e := range es {
		fmt.Println(e.Department, e.Seniority, e.Name)
	}

	// Output:
	// Engineering 5 Aaron
	// Engineering 5 Bob
	// Engineering 3 Alice
	// HR 1 Carol
}
