# GoExts Generic Utilities

[![Go Report Card](https://goreportcard.com/badge/github.com/goexts/generic)](https://goreportcard.com/report/github.com/goexts/generic)
[![GoDoc](https://godoc.org/github.com/goexts/generic?status.svg)](https://godoc.org/github.com/goexts/generic)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

A modern, robust, and type-safe collection of generic utilities for Go, designed to solve common problems with elegant, high-performance APIs. This library aims to be the foundational toolkit for modern Go development.

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

## Installation

```shell
go get github.com/goexts/generic
```

## Featured Package: `configure`

To showcase the design philosophy of this library, here is a quick look at the `configure` package. It provides a best-in-class toolset for object creation, enabling a clean separation of concerns between building a configuration object and compiling a final product.

The following example demonstrates how to create a fully configured `*http.Client` from a dedicated `ClientConfig` object:

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

// 1. Define your configuration object and its options.
type ClientConfig struct {
	Timeout   time.Duration
	Transport http.RoundTripper
}

type Option = configure.Option[ClientConfig]

func WithTimeout(d time.Duration) Option {
	return func(c *ClientConfig) { c.Timeout = d }
}

func WithTransport(rt http.RoundTripper) Option {
	return func(c *ClientConfig) { c.Transport = rt }
}

// 2. Define your factory function (the "compiler").
func NewHttpClient(c *ClientConfig) (*http.Client, error) {
	return &http.Client{
		Timeout:   c.Timeout,
		Transport: c.Transport,
	}, nil
}

func main() {
	// 3. Use the Builder to collect options, then use Compile to create the final product.
	configBuilder := configure.NewBuilder[ClientConfig]().
		Add(WithTimeout(20 * time.Second)).
		Add(WithTransport(http.DefaultTransport))

	httpClient, err := configure.Compile(configBuilder, NewHttpClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully created http.Client with timeout: %s\n", httpClient.Timeout)
}
```

## Contributing

Contributions of all kinds are welcome! Please see our Contributing Guide for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
