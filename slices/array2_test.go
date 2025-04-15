//go:build !race

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package slices_test implements the functions, types, and interfaces for the module.
package slices_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/slices"
)

func TestInsertWithMiddle(t *testing.T) {
	// Arrr! Let's put this element right in the middle of the action!
	s := []int{1, 2, 4, 5}
	v := 3
	fn := func(a, b int) bool { return a > b }
	result := slices.InsertWith(s, v, fn)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestInsertWithBeginning(t *testing.T) {
	// Ahoy! Let's sneak this element at the very start!
	s := []int{2, 3, 4, 5}
	v := 1
	fn := func(a, b int) bool { return a > b }
	result := slices.InsertWith(s, v, fn)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestInsertWithEnd(t *testing.T) {
	// Avast! We'll stash this element at the end of the line!
	s := []int{1, 2, 3, 4}
	v := 5
	fn := func(a, b int) bool { return a > b }
	result := slices.InsertWith(s, v, fn)
	expected := []int{1, 2, 3, 4, 5}
	assert.Equal(t, expected, result)
}

func TestInsertIntoNilSlice(t *testing.T) {
	// Shiver me timbers! Let's see what happens with a nil slice!
	var s []int
	v := 1
	fn := func(a, b int) bool { return a > b }
	result := slices.InsertWith(s, v, fn)
	expected := []int{1}
	assert.Equal(t, expected, result)
}

func TestInsertWithAlwaysTrueComparison(t *testing.T) {
	// Yo-ho-ho! This comparison be always true!
	s := []int{1, 2, 3}
	v := 0
	fn := func(a, b int) bool { return true }
	result := slices.InsertWith(s, v, fn)
	expected := []int{0, 1, 2, 3}
	assert.Equal(t, expected, result)
}

func TestInsertIntoSliceWithDuplicates(t *testing.T) {
	// Blimey! Let's see how it handles duplicates!
	s := []int{1, 2, 2, 3}
	v := 2
	fn := func(a, b int) bool { return a > b }
	result := slices.InsertWith(s, v, fn)
	expected := []int{1, 2, 2, 2, 3}
	assert.Equal(t, expected, result)
}
