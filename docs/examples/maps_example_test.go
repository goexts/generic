// Package examples provides examples for using the maps package, including Report Q3.
package examples

import (
	"fmt"

	"github.com/goexts/generic/maps"
	"github.com/goexts/generic/slices"
	"github.com/goexts/generic/strings"
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

// ExampleFirstKey demonstrates the usage of the FirstKey function.
func ExampleFirstKey() {
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	// Example 1: Key exists at the beginning of the list
	key, found := maps.FirstKey(m, "apple", "grape")
	fmt.Printf("FirstKey(\"apple\", \"grape\"): key=%q, found=%t\n", key, found)

	// Example 2: Key exists later in the list
	key, found = maps.FirstKey(m, "grape", "banana", "kiwi")
	fmt.Printf("FirstKey(\"grape\", \"banana\", \"kiwi\"): key=%q, found=%t\n", key, found)

	// Example 3: No key exists
	key, found = maps.FirstKey(m, "orange", "lemon")
	fmt.Printf("FirstKey(\"orange\", \"lemon\"): key=%q, found=%t\n", key, found)

	// Example 4: Empty key list
	key, found = maps.FirstKey(m)
	fmt.Printf("FirstKey(): key=%q, found=%t\n", key, found)

	// Output:
	// FirstKey("apple", "grape"): key="apple", found=true
	// FirstKey("grape", "banana", "kiwi"): key="banana", found=true
	// FirstKey("orange", "lemon"): key="", found=false
	// FirstKey(): key="", found=false
}

// ExampleFirstKeyBy demonstrates the usage of the FirstKeyBy function with a transform.
func ExampleFirstKeyBy() {
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	// Case-insensitive string comparison transform
	toLowerCase := func(s string) string {
		return strings.ToLower(s)
	}

	// Example 1: Case-insensitive match
	key, found := maps.FirstKeyBy(m, toLowerCase, "Apple", "Grape")
	fmt.Printf("FirstKeyBy(\"Apple\", \"Grape\"): key=%q, found=%t\n", key, found)

	// Example 2: Case-insensitive match later in the list
	key, found = maps.FirstKeyBy(m, toLowerCase, "Grape", "BANANA", "Kiwi")
	fmt.Printf("FirstKeyBy(\"Grape\", \"BANANA\", \"Kiwi\"): key=%q, found=%t\n", key, found)

	// Example 3: No match
	key, found = maps.FirstKeyBy(m, toLowerCase, "Orange", "Lemon")
	fmt.Printf("FirstKeyBy(\"Orange\", \"Lemon\"): key=%q, found=%t\n", key, found)

	// Example 4: Transform from int to string
	mIntKeys := map[string]string{"10": "ten", "20": "twenty"}
	intToString := func(i int) string {
		return fmt.Sprintf("%d", i)
	}
	intKey, intFound := maps.FirstKeyBy(mIntKeys, intToString, 5, 10, 15)
	fmt.Printf("FirstKeyBy(int keys): key=%v, found=%t\n", intKey, intFound)

	// Output:
	// FirstKeyBy("Apple", "Grape"): key="Apple", found=true
	// FirstKeyBy("Grape", "BANANA", "Kiwi"): key="BANANA", found=true
	// FirstKeyBy("Orange", "Lemon"): key="", found=false
	// FirstKeyBy(int keys): key=10, found=true
}
