// Package examples provides comprehensive examples for using the slices package.
// This file expands previous examples and includes solutions mapped to report.log Q1, Q2, Q7, Q8.
package examples

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

func SlicesHelpers() {
	users := []User{{1, "Alice"}, {2, "Bob"}}
	names := slices.Map(users, func(u User) string { return u.Name })
	fmt.Println("User names:", names)
}

func ExampleSlicesHelpers() {
	SlicesHelpers()

	// Output:
	// User names: [Alice Bob]
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

// ExampleGetCompletedTaskTitles demonstrates how to use GetCompletedTaskTitles and Summaries functions.
// It shows filtering completed tasks and summarizing log entries.
//
// Output:
// Completed task titles: [Write Ship]
// Summaries: [{INFO 2} {WARN 1}]
func ExampleGetCompletedTaskTitles() {
	tasks := []Task{{1, "Write", true}, {2, "Test", false}, {3, "Ship", true}}
	fmt.Println("Completed task titles:", GetCompletedTaskTitles(tasks))

	// Q8 demo
	logs := []LogEntry{{"DEBUG", "d1"}, {"INFO", "ok"}, {"WARN", "w"}}
	fmt.Println("Summaries:", Summaries(logs))
}
