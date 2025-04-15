//go:build !race

// Package slices_test implements the functions, types, and interfaces for the module.
package slices_test

import (
	"testing"

	"github.com/goexts/generic/slices"
)

func TestRemoveElementsWhereComparisonTrue(t *testing.T) {
	// Arrr, let's set sail with this test!
	input := []int{1, 2, 3, 4, 5}
	fn := func(a int) bool { return a%2 == 0 }
	result := slices.RemoveWith(input, fn)
	expected := []int{1, 3, 5}

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestReturnNewSliceWithNonMatchingElements(t *testing.T) {
	// Ahoy! Time to test the new slice!
	input := []string{"apple", "banana", "cherry"}
	fn := func(a string) bool { return a == "banana" }
	result := slices.RemoveWith(input, fn)
	expected := []string{"apple", "cherry"}

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestWorksWithDifferentTypes(t *testing.T) {
	// Yo-ho-ho! Testing with different types!
	input := []float64{1.1, 2.2, 3.3}
	fn := func(a float64) bool { return a > 2.0 }
	result := slices.RemoveWith(input, fn)
	expected := []float64{1.1}

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestHandlesEmptySlices(t *testing.T) {
	// Shiver me timbers! An empty slice!
	input := []int{}
	fn := func(a int) bool { return a > 0 }
	result := slices.RemoveWith(input, fn)
	expected := []int{}

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestAllElementsSatisfyComparisonFunction(t *testing.T) {
	// Avast! All elements be satisfying the function!
	input := []int{2, 4, 6}
	fn := func(a int) bool { return a%2 == 0 }
	result := slices.RemoveWith(input, fn)
	expected := []int{}

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestNoElementsSatisfyComparisonFunction(t *testing.T) {
	// Arrr! None of these elements satisfy the function!
	input := []int{1, 3, 5}
	fn := func(a int) bool { return a%2 == 0 }
	result := slices.RemoveWith(input, fn)
	expected := input

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestHandlesDuplicateElements(t *testing.T) {
	// Yo-ho-ho! Duplicates ahead!
	input := []int{1, 2, 2, 3}
	fn := func(a int) bool { return a == 2 }
	result := slices.RemoveWith(input, fn)
	expected := []int{1, 3}

	if !equal(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
