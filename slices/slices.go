// Package slices implements the functions, types, and interfaces for the module.
package slices

import (
	_ "golang.org/x/exp/slices" // Import for side effects (type definitions)
)

//go:generate adptool slices.go
//go:adapter:package golang.org/x/exp/slices
