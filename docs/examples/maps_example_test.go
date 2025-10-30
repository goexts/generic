// Package examples provides examples for using the maps package, including Report Q3.
package examples

import (
	"fmt"

	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/slices"
)

func Example() {
	userPermissions := map[string]string{"admin": "all", "editor": "write", "viewer": "read"}
	keys := maps.Keys(userPermissions)
	vals := maps.Values(userPermissions)
	// Sort for deterministic output (optional)
	slices.Sort(keys)
	slices.Sort(vals)
	fmt.Println("Keys:", keys)
	fmt.Println("Values:", vals)

	// Output:
	// Keys: [admin editor viewer]
	// Values: [all read write]
}
