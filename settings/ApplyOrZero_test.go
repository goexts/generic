//go:build !race

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package settings implements the functions, types, and interfaces for the module.
package settings

import (
	"testing"
)

// ApplyOrZero initializes a zero value of type S and applies functions
func TestApplyOrZeroInitializesZeroValue(t *testing.T) {
	// Arrr, let's see if it starts with a zero value!
	var modifyFunc = func(s *int) {
		*s = 42
	}
	result := ApplyOrZero(modifyFunc)
	if result == nil || *result != 42 {
		t.Errorf("Expected 42, but got %v", result)
	}
}

// Functions modify the zero value correctly when applied
func TestFunctionsModifyZeroValue(t *testing.T) {
	// Aye, let's see if the functions be doin' their job!
	var modifyFunc = func(s *int) {
		*s = 7
	}
	result := ApplyOrZero(modifyFunc)
	if result == nil || *result != 7 {
		t.Errorf("Expected 7, but got %v", result)
	}
}

// Returns a pointer to the modified zero value
func TestReturnsPointerToModifiedValue(t *testing.T) {
	// Avast! Check if it returns a pointer to the treasure!
	var modifyFunc = func(s *int) {
		*s = 99
	}
	result := ApplyOrZero(modifyFunc)
	if result == nil || *result != 99 {
		t.Errorf("Expected pointer to 99, but got %v", result)
	}
}

// No functions provided, returns a zero value of type S
func TestNoFunctionsReturnsZeroValue(t *testing.T) {
	// Shiver me timbers! No functions, no changes!
	result := ApplyOrZero[int]()
	if result == nil || *result != 0 {
		t.Errorf("Expected zero value, but got %v", result)
	}
}

// Functions that do not modify the value, returns unchanged zero value
func TestUnchangedZeroValueWithNoModification(t *testing.T) {
	// Arrr, let's see if it stays the same!
	var noOpFunc = func(s *int) {}
	result := ApplyOrZero(noOpFunc)
	if result == nil || *result != 0 {
		t.Errorf("Expected unchanged zero value, but got %v", result)
	}
}

// Supports various types for S, including complex structs
type Pirate struct {
	Name string
	Age  int
}

func TestSupportsVariousTypesIncludingStructs(t *testing.T) {
	// Yo ho ho! Let's see if it works with me pirate crew!
	var modifyFunc = func(p *Pirate) {
		p.Name = "Blackbeard"
		p.Age = 40
	}
	result := ApplyOrZero(modifyFunc)
	if result == nil || result.Name != "Blackbeard" || result.Age != 40 {
		t.Errorf("Expected Blackbeard aged 40, but got %v", result)
	}
}
