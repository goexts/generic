//go:build generate

//go:generate go run github.com/princjef/gomarkdoc/cmd/gomarkdoc --exclude-dirs ./docs/... --output docs/api/api.md ./...

// Package main provides tools for generating project documentation.
// This file is used with 'go generate' to update project documentation.
package main

// Run 'go generate' to update the documentation.
// This will generate/update the API documentation in docs/api.md
func main() {
	// This file is just a placeholder for go generate.
	// The actual documentation generation is handled by the go:generate directive above.
}
