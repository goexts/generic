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

// String returns a string representation of the Product in the format "ID:$Price"
func (p Product) String() string {
	return fmt.Sprintf("%s:$%.2f", p.ID, p.Price)
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

// GetUserNames extracts names from a slice of User structs.
// It uses generic/slices.Map for a clean, functional approach.
//
// Example:
//
//	users := []User{{1, "Alice"}, {2, "Bob"}}
//	names := GetUserNames(users) // returns ["Alice", "Bob"]
func GetUserNames(users []User) []string {
	return slices.Map(users, func(u User) string { return u.Name })
}

// ExampleGetUserNames demonstrates how to use GetUserNames function.
func ExampleGetUserNames() {
	users := []User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	names := GetUserNames(users)
	fmt.Println("User names:", names)

	// Output:
	// User names: [Alice Bob Charlie]
}

// ExampleFilterExpensiveProducts demonstrates how to use FilterExpensiveProducts function.
func ExampleFilterExpensiveProducts() {
	products := []Product{
		{ID: "A", Price: 9.99},
		{ID: "B", Price: 14.99},
		{ID: "C", Price: 4.99},
		{ID: "D", Price: 19.99},
	}

	expensive := FilterExpensiveProducts(products, 10.0)
	fmt.Println("Expensive products:", expensive)

	// Output:
	// Expensive products: [B:$14.99 D:$19.99]
}

// SlicesHelpers demonstrates various slice operations using the generic/slices package.
// This is a helper function used by ExampleSlicesHelpers.
func SlicesHelpers() {
	// Example of GetUserNames
	users := []User{{1, "Alice"}, {2, "Bob"}}
	names := GetUserNames(users)
	fmt.Println("User names:", names)

	// Example of FilterExpensiveProducts
	products := []Product{{"A", 9.99}, {"B", 14.99}, {"C", 4.99}}
	expensive := FilterExpensiveProducts(products, 10.0)
	fmt.Println("Expensive products:", expensive)

	// Example of GetCompletedTaskTitles
	tasks := []Task{{1, "Write", false}, {2, "Test", true}, {3, "Deploy", true}}
	completedTitles := GetCompletedTaskTitles(tasks)
	fmt.Println("Completed tasks:", completedTitles)
}

// ExampleSlicesHelpers demonstrates multiple slice operations in one example.
func ExampleSlicesHelpers() {
	SlicesHelpers()

	// Output:
	// User names: [Alice Bob]
	// Expensive products: [B:$14.99]
	// Completed tasks: [Test Deploy]
}

// FilterExpensiveProducts returns products with price >= minPrice.
// It uses generic/slices.Filter to create a new slice containing only the products
// that satisfy the given predicate (price >= minPrice).
func FilterExpensiveProducts(products []Product, minPrice float64) []Product {
	return slices.Filter(products, func(p Product) bool { return p.Price >= minPrice })
}

// GetCompletedTaskTitles returns the titles of all completed tasks.
// It demonstrates chaining of Filter and Map operations.
func GetCompletedTaskTitles(tasks []Task) []string {
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
