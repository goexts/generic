// Package main provides comprehensive examples for using the slices package.
// This file expands previous examples and includes solutions mapped to report.log Q1, Q2, Q7, Q8.
package main

import (
	"fmt"

	"github.com/goexts/generic/slices"
)

type User struct {
	ID   int
	Name string
}
type Product struct {
	ID    string
	Price float64
}
type Task struct {
	ID        int
	Title     string
	Completed bool
}
type LogEntry struct {
	Level string
	Msg   string
}
type Summary struct {
	Level string
	Len   int
}

func GetUserNames(users []User) []string { // Q1
	return slices.Map(users, func(u User) string { return u.Name })
}

func FilterExpensiveProducts(products []Product, minPrice float64) []Product { // Q2
	return slices.Filter(products, func(p Product) bool { return p.Price >= minPrice })
}

func GetCompletedTaskTitles(tasks []Task) []string { // Q7
	return slices.Map(
		slices.Filter(tasks, func(t Task) bool { return t.Completed }),
		func(t Task) string { return t.Title },
	)
}

// Summaries demonstrates a performance-conscious pipeline (Q8).
// Strategy: filter first to minimize size, then map. Optionally preallocate.
func Summaries(entries []LogEntry) []Summary {
	filtered := slices.Filter(entries, func(e LogEntry) bool { return e.Level != "DEBUG" })
	// Preallocate to reduce allocations.
	res := make([]Summary, 0, len(filtered))
	for _, e := range filtered {
		res = append(res, Summary{Level: e.Level, Len: len(e.Msg)})
	}
	return res
}

func main() {
	// Example: Basic Map/Filter/Reduce
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubled := slices.Map(nums, func(x int) int { return x * 2 })
	fmt.Println("Doubled:", doubled)
	evens := slices.Filter(doubled, func(x int) bool { return x%2 == 0 })
	fmt.Println("Even numbers:", evens)
	sum := slices.Reduce(nums, 0, func(acc, x int) int { return acc + x })
	fmt.Println("Sum:", sum)

	// Example: Strings transform
	names := []string{"alice", "bob", "charlie", "dave"}
	capitalized := slices.Map(names, func(s string) string {
		if len(s) == 0 {
			return s
		}
		return string(append([]byte{s[0] - 32}, s[1:]...))
	})
	fmt.Println("Capitalized:", capitalized)

	// Example: FilterIncluded / FilterExcluded
	allNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evensOnly := slices.FilterIncluded(allNumbers, []int{2, 4, 6, 8, 10})
	fmt.Println("Even numbers (FilterIncluded):", evensOnly)
	oddsOnly := slices.FilterExcluded(allNumbers, []int{2, 4, 6, 8, 10})
	fmt.Println("Odd numbers (FilterExcluded):", oddsOnly)

	// Q1 demo
	users := []User{{1, "Alice"}, {2, "Bob"}}
	fmt.Println("User names:", GetUserNames(users))

	// Q2 demo
	products := []Product{{"A", 10.0}, {"B", 99.5}, {"C", 100.0}}
	fmt.Println("Expensive products:", FilterExpensiveProducts(products, 99.5))

	// Q7 demo
	tasks := []Task{{1, "Write", true}, {2, "Test", false}, {3, "Ship", true}}
	fmt.Println("Completed task titles:", GetCompletedTaskTitles(tasks))

	// Q8 demo
	logs := []LogEntry{{"DEBUG", "d1"}, {"INFO", "ok"}, {"WARN", "w"}}
	fmt.Println("Summaries:", Summaries(logs))
}
