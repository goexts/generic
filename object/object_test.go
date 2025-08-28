package object

import (
	"fmt"
	"testing"
)

// Person is an example struct that embeds BaseObject.
type Person struct {
	BaseObject // Embedding BaseObject to get default Object methods.
	Name       string
	Age        int
}

// NewPerson creates a new Person object.
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// String provides a custom string representation for Person.
func (p *Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// Equals provides a custom equality check for Person.
func (p *Person) Equals(other Object) bool {
	if other == nil {
		return false
	}
	if p2, ok := other.(*Person); ok {
		return p.Name == p2.Name && p.Age == p2.Age
	}
	return false
}

// HashCode provides a custom hash code for Person.
func (p *Person) HashCode() int {
	result := 1
	result = 31*result + len(p.Name)
	result = 31*result + p.Age
	return result
}

func TestObject(t *testing.T) {
	// Create two person objects with the same values.
	p1 := NewPerson("Alice", 30)
	p2 := NewPerson("Alice", 30)
	// Create a different person object.
	p3 := NewPerson("Bob", 25)

	// Test String() method
	fmt.Println("p1:", p1.String())
	fmt.Println("p2:", p2.String())
	fmt.Println("p3:", p3.String())

	// Test Equals() method
	fmt.Println("p1 equals p2?", p1.Equals(p2)) // Should be true
	fmt.Println("p1 equals p3?", p1.Equals(p3)) // Should be false

	// Test HashCode() method
	fmt.Println("p1 HashCode:", p1.HashCode())
	fmt.Println("p2 HashCode:", p2.HashCode()) // Should be same as p1
	fmt.Println("p3 HashCode:", p3.HashCode())

	// Using Object interface
	var obj Object = p1
	fmt.Println("Object variable:", obj.String())
}
