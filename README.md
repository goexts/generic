# GoExts Generic Utilities

<div align="right">
  <a href="./README.md">English</a> | 
  <a href="./README.zh-CN.md">中文</a> | 
  <a href="./README.ja-JP.md">日本語</a>
</div>

[![Go Report Card](https://goreportcard.com/badge/github.com/goexts/generic)](https://goreportcard.com/report/github.com/goexts/generic)
[![GoDoc](https://godoc.org/github.com/goexts/generic?status.svg)](https://pkg.go.dev/github.com/goexts/generic)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/badge/release-MIT-blue.svg)](https://github.com/goexts/generic/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/goexts/generic)](go.mod)
[![GitHub stars](https://img.shields.io/github/stars/goexts/generic?style=social)](https://github.com/goexts/generic/stargazers)

A modern, robust, and type-safe collection of generic utilities for Go, designed to solve common problems with elegant, high-performance APIs.

> **Status**: **Stable** - This project is production-ready and follows semantic versioning.

## Installation

```bash
go get github.com/goexts/generic@latest
```

## Documentation

For a complete guide, API reference, and usage examples for all packages, please visit the official Go documentation:

**[pkg.go.dev/github.com/goexts/generic](https://pkg.go.dev/github.go.dev/github.com/goexts/generic)**

## Packages

This library provides a rich set of independent, generic packages:

*   **`cast`**: Safe, generic type-casting functions.
*   **`cmp`**: Generic comparison functions for sorting and ordering.
*   **`cond`**: Ternary-like conditional functions.
*   **`configure`**: A powerful implementation of the Functional Options Pattern.
*   **`maps`**: A suite of generic functions for common map operations (adapter for `x/exp/maps`).
*   **`must`**: Panic-on-error wrappers for cleaner initialization code.
*   **`promise`**: A generic, JavaScript-like `Promise` implementation for managing asynchronous operations.
*   **`ptr`**: Helper functions for creating pointers from literal values.
*   **`res`**: A generic, Rust-inspired `Result[T]` type for expressive error handling.
*   **`set`**: Stateless, slice-based set operations.
*   **`slices`**: A comprehensive suite of functions for slice operations.
*   **`strings`**: A collection of functions for string manipulation (adapter for the standard `strings` package).

## Featured Example: `configure`

The `configure` package provides a robust and type-safe implementation of the **Functional Options Pattern**. This pattern is ideal for creating complex objects with multiple optional parameters in a clean and readable way.

### 1. Basic Functional Options (`configure.Apply`)

This demonstrates the core pattern: defining options and applying them to a default configuration. Note that `configure.Apply` expects a slice of options.

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

// ClientConfig holds the configuration for our http.Client.
type ClientConfig struct {
	Timeout   time.Duration
	Transport http.RoundTripper
}

// Option defines the functional option type for our configuration.
type Option = configure.Option[ClientConfig]

// WithTimeout returns an option to set the client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *ClientConfig) {
		c.Timeout = d
	}
}

// WithTransport returns an option to set the client's transport.
func WithTransport(rt http.RoundTripper) Option {
	return func(c *ClientConfig) {
		c.Transport = rt
	}
}

// NewClient creates a new http.Client with default settings,
// then applies the provided options.
func NewClient(opts ...Option) *http.Client {
	// 1. Start with a default configuration.
	config := &ClientConfig{
		Timeout:   10 * time.Second,
		Transport: http.DefaultTransport,
	}

	// 2. Apply any user-provided options over the defaults.
	// configure.Apply modifies the config in-place and returns it.
	configure.Apply(config, opts)

	// 3. The final, configured object is now ready to be used.
	return &http.Client{
		Timeout:   config.Timeout,
		Transport: config.Transport,
	}
}

func main() {
	// Create a client with default settings (no options).
	defaultClient := NewClient()
	fmt.Printf("Default client timeout: %s\n", defaultClient.Timeout)

	// Create a client with a custom timeout, overriding the default.
	// Pass options as variadic arguments.
	customClient := NewClient(WithTimeout(30 * time.Second))
	fmt.Printf("Custom client timeout: %s\n", customClient.Timeout)

	// If you have individual options, you can pass them directly:
	// anotherClient := NewClient(WithTimeout(20 * time.Second), WithTransport(&http.Transport{}))
}
```

### 2. Advanced Configuration with `Builder`

The `Builder` provides a fluent interface for collecting options incrementally, which is useful when options are gathered from various sources or in multiple stages. It also supports setting a base configuration.

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goexts/generic/configure"
)

// ClientConfig holds the configuration for our http.Client.
type ClientConfig struct {
	Timeout       time.Duration
	Transport     http.RoundTripper
	EnableTracing bool // Added to demonstrate AddWhen's optIfFalse
}

// Option defines the functional option type for our configuration.
type Option = configure.Option[ClientConfig]

// WithTimeout returns an option to set the client timeout.
func WithTimeout(d time.Duration) Option {
	return func(c *ClientConfig) {
		c.Timeout = d
	}
}

// WithTransport returns an option to set the client's transport.
func WithTransport(rt http.RoundTripper) Option {
	return func(c *ClientConfig) {
		c.Transport = rt
	}
}

// WithTracing enables or disables tracing.
func WithTracing(enable bool) Option {
	return func(c *ClientConfig) {
		c.EnableTracing = enable
	}
}

// NewClient is the factory function for Compile.
// It takes the final, fully-built configuration and creates the product (*http.Client).
func NewClient(config *ClientConfig) (*http.Client, error) {
	// In a real scenario, you might use config.EnableTracing here
	// to configure the http.Client or a wrapper around it.
	return &http.Client{
		Timeout:   config.Timeout,
		Transport: config.Transport,
	}, nil
}

func main() {
	// Define a base configuration that can be reused.
	baseConfig := &ClientConfig{
		Timeout:       5 * time.Second,
		Transport:     http.DefaultTransport,
		EnableTracing: false,
	}

	// Create a builder, passing the base configuration directly to NewBuilder.
	builder := configure.NewBuilder[ClientConfig](baseConfig).
		Add(WithTimeout(15 * time.Second)). // Overrides baseConfig.Timeout
		// Use AddWhen with optIfTrue and optIfFalse
		AddWhen(true, WithTracing(true), WithTracing(false))

	// Compile the final product using the builder and a factory function.
	// Note: configure.Compile now expects factory first, then builder.
	client1, err := configure.Compile(NewClient, builder)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client 1 (Builder) timeout: %s\n", client1.Timeout)

	// Demonstrate using Chain to group options before adding to builder
	commonOptions := configure.Chain(
		WithTransport(&http.Transport{}), // Custom transport
		WithTimeout(20 * time.Second),
	)

	client2, err := configure.Compile(
		NewClient,
		configure.NewBuilder[ClientConfig](baseConfig).
			Add(commonOptions). // Add chained options
			// AddWhen with false condition, so optIfFalse (WithTracing(false)) will be applied
			AddWhen(false, WithTracing(true), WithTracing(false)),
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Client 2 (Builder) transport type: %T, timeout: %s\n", client2.Transport, client2.Timeout)

	// Example of using Builder directly as an option (implements ApplierE)
	// This is useful if you want to apply a set of options defined by a builder
	// to an existing config object or within another ApplyAny call.
	existingConfig := &ClientConfig{Timeout: 1 * time.Second}
	err = configure.NewBuilder[ClientConfig]().
		Add(WithTimeout(30 * time.Second)).
		Apply(existingConfig) // Apply builder's options to an existing config
	if err != nil {
		panic(err)
	}
	fmt.Printf("Existing config after builder.Apply: %s\n", existingConfig.Timeout)
}
```

## Contributing

We welcome contributions! Please see our [Contributing Guide](.github/CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
