// Copy elements from source to target array correctly
package slices_test

import (
	"testing"

	"github.com/goexts/generic/slices"
)

func TestCopyElementsCorrectly(t *testing.T) {
	source := []int{1, 2, 3, 4, 5}
	target := []int{6, 7, 8}
	result := slices.CopyAt(source, target, 2)

	expected := []int{1, 2, 6, 7, 8}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Arrr! Expected %d at index %d but got %d", v, i, result[i])
		}
	}
}

func TestTargetFitsWithinSourceCapacity(t *testing.T) {
	source := make([]int, 5, 10)
	target := []int{6, 7, 8}
	result := slices.CopyAt(source, target, 2)

	expected := []int{0, 0, 6, 7, 8}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Shiver me timbers! Expected %d at index %d but got %d", v, i, result[i])
		}
	}
}

func TestReturnModifiedTargetArray(t *testing.T) {
	source := []int{1, 2, 3}
	target := []int{4, 5}
	result := slices.CopyAt(source, target, 1)

	expected := []int{1, 4, 5}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Blimey! Expected %d at index %d but got %d", v, i, result[i])
		}
	}
}

func TestSourceInsufficientCapacity(t *testing.T) {
	source := make([]int, 3)
	target := []int{4, 5}
	result := slices.CopyAt(source, target, 2)

	expected := []int{0, 0, 4, 5}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Ahoy! Expected %d at index %d but got %d", v, i, result[i])
		}
	}
}

func TestNegativeIndexValues(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Walk the plank! Expected panic for negative index but did not get one")
		}
	}()

	source := []int{1, 2, 3}
	target := []int{4, 5}
	_ = slices.CopyAt(source, target, -1)
}

func TestIndexGreaterThanSourceLength(t *testing.T) {
	source := []int{1, 2}
	target := []int{3}
	result := slices.CopyAt(source, target, 3)

	expected := []int{1, 2, 0, 3}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Yo-ho-ho! Expected %d at index %d but got %d", v, i, result[i])
		}
	}
}

func TestEmptySourceAndTargetArrays(t *testing.T) {
	source := []int{}
	target := []int{}
	result := slices.CopyAt(source, target, 0)

	if len(result) != 0 {
		t.Errorf("Avast! Expected empty array but got length %d", len(source))
	}
}
