# GoExts Generic Utilities

[![Go Report Card](https://goreportcard.com/badge/github.com/goexts/generic)](https://goreportcard.com/report/github.com/goexts/generic)
[![GoDoc](https://godoc.org/github.com/goexts/generic?status.svg)](https://godoc.org/github.com/goexts/generic)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/goexts/generic.svg)](https://github.com/goexts/generic/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/goexts/generic)](go.mod)
[![GitHub stars](https://img.shields.io/github/stars/goexts/generic?style=social)](https://github.com/goexts/generic/stargazers)

A modern, robust, and type-safe collection of generic utilities for Go, designed to solve common problems with elegant, high-performance APIs.

> **Status**: **Stable** - This project is production-ready and follows semantic versioning.

## Installation

```bash
go get github.com/goexts/generic@latest
```

## Documentation

For a complete guide, API reference, and usage examples, please visit the official Go documentation:

**[pkg.go.dev/github.com/goexts/generic](https://pkg.go.dev/github.com/goexts/generic)**

## Packages

This library provides a rich set of independent, generic packages:

*   **`cast`**: Safe, generic type-casting functions (`As`, `AsE`).
*   **`cmp`**: Generic comparison functions (`Compare`, `Min`, `Max`) for sorting and ordering.
*   **`cond`**: Ternary-like conditional functions (`If`, `IfFunc`).
*   **`configure`**: A powerful implementation of the Functional Options Pattern.
*   **`maps`**: A suite of generic functions for common map operations (adapter for `x/exp/maps`).
*   **`must`**: Panic-on-error wrappers (`Must`) for cleaner initialization code.
*   **`promise`**: A generic, JavaScript-like `Promise` implementation for managing asynchronous operations.
*   **`ptr`**: Helper functions for creating pointers from literal values (`Of`, `Value`).
*   **`res`**: A generic, Rust-inspired `Result[T]` type for expressive error handling.
*   **`set`**: Stateless, slice-based set operations (`Union`, `Intersection`, `Difference`).
*   **`slices`**: A comprehensive suite of functions for slice operations (adapter for `x/exp/slices`).
    *   **`slices/bytes`**: Adapter for the standard `bytes` package.
    *   **`slices/runes`**: Adapter for the `x/text/runes` package.
*   **`strings`**: A collection of functions for string manipulation (adapter for the standard `strings` package).

## Contributing

We welcome all contributions! Please read our [Contributing Guide](.github/CONTRIBUTING.md) and [Code of Conduct](.github/CODE_OF_CONDUCT.md).

## Community

- **Issues**: [GitHub Issues](https://github.com/goexts/generic/issues)
- **Discussions**: [GitHub Discussions](https://github.com/goexts/generic/discussions)

## License

This project is licensed under the MIT License. See the LICENSE file for details.
