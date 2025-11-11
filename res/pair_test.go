package res

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPair(t *testing.T) {
	t.Run("NewPair and Values", func(t *testing.T) {
		p := NewPair(42, "hello")
		first, second := p.Values()
		assert.Equal(t, 42, first)
		assert.Equal(t, "hello", second)
	})

	t.Run("First and Second methods", func(t *testing.T) {
		p := NewPair("test", 123)
		assert.Equal(t, "test", p.First())
		assert.Equal(t, 123, p.Second())
	})

	t.Run("Map", func(t *testing.T) {
		p := NewPair(10, 20.0)
		mapped := p.Map(
			func(x int) int { return x * 2 },
			func(y float64) float64 { return y / 2 },
		)
		first, second := mapped.Values()
		assert.Equal(t, 20, first)
		assert.Equal(t, 10.0, second)
	})

	t.Run("MapFirst", func(t *testing.T) {
		p := NewPair("hello", 42)
		mapped := p.MapFirst(func(s string) string { return s + " world" })
		first, second := mapped.Values()
		assert.Equal(t, "hello world", first)
		assert.Equal(t, 42, second)
	})

	t.Run("MapSecond", func(t *testing.T) {
		p := NewPair(3.14, true)
		mapped := p.MapSecond(func(b bool) bool { return !b })
		first, second := mapped.Values()
		assert.Equal(t, 3.14, first)
		assert.False(t, second)
	})

	t.Run("Swap", func(t *testing.T) {
		p := NewPair(42, "hello")
		swapped := p.Swap()
		first, second := swapped.Values()
		assert.Equal(t, "hello", first)
		assert.Equal(t, 42, second)
	})

	t.Run("WithFirst", func(t *testing.T) {
		p := NewPair(1, "one")
		newP := p.WithFirst(2)
		assert.Equal(t, 2, newP.First())
		assert.Equal(t, "one", newP.Second())
		// Original pair should remain unchanged
		assert.Equal(t, 1, p.First())
	})

	t.Run("WithSecond", func(t *testing.T) {
		p := NewPair(1, "one")
		newP := p.WithSecond("two")
		assert.Equal(t, 1, newP.First())
		assert.Equal(t, "two", newP.Second())
		// Original pair should remain unchanged
		assert.Equal(t, "one", p.Second())
	})

	t.Run("PairWithDifferentTypes", func(t *testing.T) {
		// Test with complex types
		type Person struct{ Name string }
		type Address struct{ City string }

		p := NewPair(
			Person{Name: "Alice"},
			Address{City: "Wonderland"},
		)

		person, address := p.Values()
		assert.Equal(t, "Alice", person.Name)
		assert.Equal(t, "Wonderland", address.City)

		// Test mapping with different types
		mapped := p.Map(
			func(p Person) Person { return Person{Name: p.Name + " (mapped)"} },
			func(a Address) Address { return Address{City: a.City + " (mapped)"} },
		)
		mappedPerson, mappedAddress := mapped.Values()
		assert.Equal(t, "Alice (mapped)", mappedPerson.Name)
		assert.Equal(t, "Wonderland (mapped)", mappedAddress.City)

		// Test Swap with different types
		swapped := p.Swap()
		swappedAddress, swappedPerson := swapped.Values()
		assert.Equal(t, "Wonderland", swappedAddress.City)
		assert.Equal(t, "Alice", swappedPerson.Name)
	})
}
