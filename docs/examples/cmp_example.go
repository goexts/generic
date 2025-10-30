// Package main shows multi-level sort using cmp with slices.SortFunc (Report Q9).
package main

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

func SortEmployees(employees []Employee) {
	slices.SortFunc(employees, func(a, b Employee) int {
		if res := cmp.Compare(a.Department, b.Department); res != 0 {
			return res
		}
		if res := cmp.Compare(b.Seniority, a.Seniority); res != 0 {
			return res
		} // desc
		return cmp.Compare(a.Name, b.Name)
	})
}

func main() {
	es := []Employee{
		{"Engineering", 3, "Alice"},
		{"Engineering", 5, "Bob"},
		{"HR", 1, "Carol"},
		{"Engineering", 5, "Aaron"},
	}
	SortEmployees(es)
	fmt.Println(es)
}
