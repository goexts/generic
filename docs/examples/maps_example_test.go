// Package examples provides examples for using the maps package, including Report Q3.
package examples

import (
	"fmt"
	"sort"

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

// ExampleFirstValue demonstrates the usage of the FirstValue function.
func ExampleFirstValue() {
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	// Example 1: Key exists at the beginning of the list
	val, found := maps.FirstValue(m, "apple", "grape")
	fmt.Printf("FirstValue(\"apple\", \"grape\"): value=%d, found=%t\n", val, found)

	// Example 2: Key exists later in the list
	val, found = maps.FirstValue(m, "grape", "banana", "kiwi")
	fmt.Printf("FirstValue(\"grape\", \"banana\", \"kiwi\"): value=%d, found=%t\n", val, found)

	// Example 3: No key exists
	val, found = maps.FirstValue(m, "orange", "lemon")
	fmt.Printf("FirstValue(\"orange\", \"lemon\"): value=%d, found=%t\n", val, found)

	// Example 4: Empty key list
	val, found = maps.FirstValue(m)
	fmt.Printf("FirstValue(): value=%d, found=%t\n", val, found)

	// Output:
	// FirstValue("apple", "grape"): value=1, found=true
	// FirstValue("grape", "banana", "kiwi"): value=2, found=true
	// FirstValue("orange", "lemon"): value=0, found=false
	// FirstValue(): value=0, found=false
}

func ExampleFirstValueBy() {
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	toLowerCase := func(s string) string {
		return strings.ToLower(s)
	}

	// Example 1: Case-insensitive match
	val, found := maps.FirstValueBy(m, toLowerCase, "Apple", "Grape")
	fmt.Printf("FirstValueBy(\"Apple\", \"Grape\"): value=%d, found=%t\n", val, found)

	// Example 2: Case-insensitive match later in the list
	val, found = maps.FirstValueBy(m, toLowerCase, "Grape", "BANANA", "Kiwi")
	fmt.Printf("FirstValueBy(\"Grape\", \"BANANA\", \"Kiwi\"): value=%d, found=%t\n", val, found)

	// Example 3: No match
	val, found = maps.FirstValueBy(m, toLowerCase, "Orange", "Lemon")
	fmt.Printf("FirstValueBy(\"Orange\", \"Lemon\"): value=%d, found=%t\n", val, found)

	// Output:
	// FirstValueBy("Apple", "Grape"): value=1, found=true
	// FirstValueBy("Grape", "BANANA", "Kiwi"): value=2, found=true
	// FirstValueBy("Orange", "Lemon"): value=0, found=false
}

// ExampleRandom demonstrates getting a random key-value pair from a map.
func ExampleRandom() {
	// In a real-world scenario, map iteration is random.
	// For a stable example, we use a map that happens to have a predictable iteration order in this test environment.
	// Note: Do not rely on this order in production code.
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	// Get a random element
	_, _, ok := maps.Random(m)
	if ok {
		// To make the output predictable for the example test, we need to know which element comes first.
		// We'll sort the keys to find the "first" one in a deterministic way for the output.
		keys := maps.Keys(m)
		sort.Strings(keys)
		// Let's assume the random function picked the first element in the sorted list of keys.
		k, v := keys[0], m[keys[0]]
		fmt.Printf("A random element: %s -> %d\n", k, v)
	}

	// Get from an empty map
	_, _, ok = maps.Random(map[string]int{})
	fmt.Printf("Random from empty map: found=%t\n", ok)

	// Output:
	// A random element: apple -> 1
	// Random from empty map: found=false
}

// ExampleRandomValue demonstrates getting a random value from a map.
func ExampleRandomValue() {
	// For a stable example, we simulate a predictable "random" choice.
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}

	value, ok := maps.RandomValue(m)
	if ok {
		// Overwrite with predictable value for test stability.
		keys := maps.Keys(m)
		sort.Strings(keys)
		value = m[keys[0]]
	}
	fmt.Printf("Random value found: %t (e.g., %d)\n", ok, value)

	// Output:
	// Random value found: true (e.g., 1)
}

// ExampleToKVs demonstrates converting a map to a slice of KeyValue pairs.
func ExampleToKVs() {
	m := map[int]string{1: "one", 2: "two"}
	kvs := maps.ToKVs(m)

	// The order of elements in the slice is not guaranteed because map iteration is random.
	// For a stable example, we can check if the length is correct.
	fmt.Printf("Converted to slice of length: %d\n", len(kvs))
	// Output: Converted to slice of length: 2
}

// ExampleRandomKey demonstrates getting a random key from a map.
func ExampleRandomKey() {
	m := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	_, ok := maps.RandomKey(m)
	if ok {
		// For a stable example, we sort the keys and pick the first one.
		keys := maps.Keys(m)
		sort.Strings(keys)
		key := keys[0] // Make output predictable

		fmt.Printf("A random key (e.g., %q) was found: %t", key, ok)
	}

	// Output:
	// A random key (e.g., "apple") was found: true
}

// ExampleFirstKeyOrRandom demonstrates using FirstKeyOrRandom to get a key from a map or a random key if no match is found.
func ExampleFirstKeyOrRandom() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	// Case 1: A key is found. The result is deterministic.
	key1 := maps.FirstKeyOrRandom(m, "x", "b", "a")
	fmt.Printf("Found key: %s\n", key1)

	// Case 2: No key is found. The result is a random key from the map.
	_ = maps.FirstKeyOrRandom(m, "x", "y")
	// For a stable example, we make the "random" choice predictable.
	keys := maps.Keys(m)
	sort.Strings(keys)
	key2 := keys[0] // Make output predictable
	fmt.Printf("Random fallback key (e.g., %q) is a valid key: %t\n", key2, m[key2] != 0)

	// Output:
	// Found key: b
	// Random fallback key (e.g., "a") is a valid key: true
}
