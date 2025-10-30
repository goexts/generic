# GoExts Generic Utilities

[![Go Report Card](https://goreportcard.com/badge/github.com/goexts/generic)](https://goreportcard.com/report/github.com/goexts/generic)
[![GoDoc](https://godoc.org/github.com/goexts/generic?status.svg)](https://godoc.org/github.com/goexts/generic)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/goexts/generic.svg)](https://github.com/goexts/generic/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/goexts/generic)](go.mod)
[![GitHub stars](https://img.shields.io/github/stars/goexts/generic?style=social)](https://github.com/goexts/generic/stargazers)

A modern, robust, and type-safe collection of generic utilities for Go, designed to solve common problems with elegant, high-performance APIs. This library aims to be the foundational toolkit for modern Go development.

> **Status**: **Stable** - This project is production-ready and follows semantic versioning.

## Features

- 🚀 **Type-Safe**: Built with Go 1.18+ generics for type safety
- ⚡ **High Performance**: Optimized for performance with zero or minimal allocations
- 🧩 **Modular**: Independent packages that can be used separately
- 🛠 **Well-Tested**: Comprehensive test coverage
- 📚 **Well-Documented**: Complete API documentation and examples

## Quick Start

### Installation

```bash
go get github.com/goexts/generic@latest
```

### Basic Usage

```go
package main

import (
	"fmt"
	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/slices"
)

func main() {
	// Working with slices
	nums := []int{1, 2, 3, 4, 5}
	doubled := slices.Map(nums, func(x int) int { return x * 2 })
	filtered := slices.Filter(doubled, func(x int) bool { return x > 5 })
	
	fmt.Println("Original:", nums)      // [1 2 3 4 5]
	fmt.Println("Doubled:", doubled)    // [2 4 6 8 10]
	fmt.Println("Filtered:", filtered)  // [6 8 10]

	// Working with maps
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	keys := maps.Keys(m)  // [a b c] (order not guaranteed)
	values := maps.Values(m)  // [1 2 3] (order not guaranteed)
	
	fmt.Println("Keys:", keys)      // [a b c]
	fmt.Println("Values:", values)  // [1 2 3]
}
```

## Documentation

For complete documentation, please visit:

- [API Documentation (pkg.go.dev)](https://pkg.go.dev/github.com/goexts/generic)
- [API Reference (gomarkdoc, auto-generated)](docs/api/api.md)
- [Examples](docs/examples/)
- [Contributing Guide](.github/CONTRIBUTING.md)

## Core Packages

This project provides a rich set of independent, generic packages:

*   **`cast`**: Provides safe, generic type-casting functions.
*   **`cmp`**: Generic comparison functions for sorting and ordering complex types.
*   **`cond`**: Functions for conditional (ternary-like) operations.
*   **`configure`**: A powerful, production-grade implementation of the Functional Options Pattern, with advanced support for compilation workflows.
*   **`maps`**: A suite of generic functions for common map operations (e.g., `Keys`, `Values`, `Clone`).
*   **`must`**: Panic-on-error wrappers (`must.Must`) for cleaner code in contexts where an error is considered a fatal condition (e.g., during initialization).
*   **`promise`**: A generic, JavaScript-like `Promise` implementation for managing asynchronous operations.
*   **`ptr`**: Helper functions (`ptr.To`) for creating pointers from literal values.
*   **`res`**: A generic, Rust-inspired `Result[T, E]` type for expressive, explicit error handling.
*   **`set`**: A generic `Set` data structure implementation with common set operations.
*   **`slices`**: A comprehensive suite of generic functions for common slice operations, with specialized sub-packages for `bytes` and `runes`.
*   **`strings`**: Generic utilities for string manipulation and conversion.

## Examples Index

- Slices basics and chaining: [docs/examples/slices_example.go](docs/examples/slices_example.go)
- More examples are documented inside each package’s `doc.go` and will appear on pkg.go.dev.

### Selected Examples

#### 1. Chaining Slice Operations (Filter & Map)

Cleanly chain functions from the `slices` package to create powerful data processing pipelines.

```go
import "github.com/goexts/generic/slices"

type Task struct {
	Title     string
	Completed bool
}

func GetCompletedTaskTitles(tasks []Task) []string {
	return slices.Map(
		slices.Filter(tasks, func(t Task) bool { return t.Completed }),
		func(t Task) string { return t.Title },
	)
}
```

#### 2. Safe Type Casting with `cast.As`

Use `cast.As` to safely handle variables of type `any` for event handlers or plugin systems.

```go
import "github.com/goexts/generic/cast"

type UserCreatedEvent struct{ UserID int }
type OrderPlacedEvent struct{ OrderID string }

func HandleEvent(event any) {
	if e, ok := cast.As[UserCreatedEvent](event); ok { /* processUserCreated(e) */ return }
	if e, ok := cast.As[OrderPlacedEvent](event); ok { /* processOrderPlaced(e) */ return }
	// Ignore other event types
}
```

#### 3. Multi-Level Sorting (Stable)

```go
import "sort"

type Employee struct{ Department string; Seniority int; Name string }

func SortEmployees(employees []Employee) {
	sort.SliceStable(employees, func(i, j int) bool {
		if employees[i].Department != employees[j].Department {
			return employees[i].Department < employees[j].Department
		}
		if employees[i].Seniority != employees[j].Seniority {
			return employees[i].Seniority > employees[j].Seniority
		}
		return employees[i].Name < employees[j].Name
	})
}
```

#### 4. Conditional Logic with `cond.If`

```go
import "github.com/goexts/generic/cond"

func GetStatusMessage(err error) string {
	return cond.If(err == nil, "Status: OK", "Status: Failed")
}
```

## Featured Package: `configure`

The `configure` package provides a robust toolset for object creation and option application.

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

type ClientConfig struct {
	Timeout   time.Duration
	Transport http.RoundTripper
}

type Option = configure.Option[ClientConfig]

func WithTimeout(d time.Duration) Option { return func(c *ClientConfig) { c.Timeout = d } }
func WithTransport(rt http.RoundTripper) Option { return func(c *ClientConfig) { c.Transport = rt } }

func NewHttpClient(c *ClientConfig) (*http.Client, error) { return &http.Client{Timeout: c.Timeout, Transport: c.Transport}, nil }

func main() {
	configBuilder := configure.NewBuilder[ClientConfig]().
		Add(WithTimeout(20 * time.Second)).
		Add(WithTransport(http.DefaultTransport))

	httpClient, err := configure.Compile(configBuilder, NewHttpClient)
	if err != nil { panic(err) }
	fmt.Printf("Successfully created http.Client with timeout: %s\n", httpClient.Timeout)
}
```

## Contributing

We welcome all contributions! Please read our [Contributing Guide](.github/CONTRIBUTING.md) and [Code of Conduct](.github/CODE_OF_CONDUCT.md).

### Development Setup

1. Fork the repository
2. Clone the repository: `git clone https://github.com/goexts/generic.git`
3. Run tests: `go test ./...`
4. Make your changes and submit a pull request

## Community

- **Issues**: [GitHub Issues](https://github.com/goexts/generic/issues)
- **Discussions**: [GitHub Discussions](https://github.com/goexts/generic/discussions)

## License

This project is licensed under the MIT License. See the LICENSE file for details.
