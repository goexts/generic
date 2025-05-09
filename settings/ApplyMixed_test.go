//go:build !race

/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package settings implements the functions, types, and interfaces for the module.
package settings

import (
	"testing"
)

// Applies a list of Func or func to a struct
func TestApplyMixedWithApplyFuncAndFunc(t *testing.T) {
	// Arrr, let's set sail with some settings!
	type Ship struct {
		Name string
	}
	s := &Ship{}
	settings := []interface{}{
		Func[Ship](func(s *Ship) { s.Name = "Black Pearl" }),
		func(s *Ship) { s.Name += " - The Fastest" },
	}
	ApplyMixed(s, settings)
	if s.Name != "Black Pearl - The Fastest" {
		t.Errorf("Expected 'Black Pearl - The Fastest', but got %s", s.Name)
	}
}

// Returns the modified struct after applying settings
func TestApplyMixedReturnsModifiedStruct(t *testing.T) {
	// Aye, the ship be modified after the storm!
	type Ship struct {
		Speed int
	}
	s := &Ship{Speed: 10}
	settings := []interface{}{
		func(s *Ship) { s.Speed += 5 },
	}
	result, _ := ApplyMixed(s, settings)
	if result.Speed != 15 {
		t.Errorf("Expected speed 15, but got %d", result.Speed)
	}
}

type Ship struct {
	Crew int
}

type CrewSetting struct {
}

func (c CrewSetting) Apply(s *Ship) {
	s.Crew = 100
}

// Handles SettingApplier interface implementations correctly
func TestApplyMixedWithApplySettingInterface(t *testing.T) {
	s := &Ship{}
	settings := []interface{}{
		CrewSetting{},
	}
	ApplyMixed(s, settings)
	if s.Crew != 100 {
		t.Errorf("Expected crew 100, but got %d", s.Crew)
	}
}

// Returns nil if the input struct pointer is nil
func TestApplyMixedWithNilStruct(t *testing.T) {
	// Shiver me timbers! A nil ship be no ship at all!
	type Ship struct {
		Name string
	}
	var s *Ship = nil
	settings := []interface{}{
		func(s *Ship) { s.Name = "Flying Dutchman" },
	}
	result, _ := ApplyMixed(s, settings)
	if result != nil {
		t.Error("Expected nil, but got a non-nil result")
	}
}

// Panics when an invalid setting type is provided
func TestApplyStrictPanicsOnInvalidType(t *testing.T) {
	// Arrr! Beware the kraken of invalid types!
	type Ship struct {
		Name string
	}
	s := &Ship{}
	settings := []interface{}{
		"Invalid Code",
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but did not panic")
		}
	}()
	ApplyStrict(s, settings)
}

// Handles an empty list of settings without error
func TestApplyMixedWithEmptySettingsList(t *testing.T) {
	// Aye, an empty list be as harmless as a calm sea!
	type Ship struct {
		Name string
	}
	s := &Ship{Name: "Silent Mary"}
	settings := []interface{}{}
	result, _ := ApplyMixed(s, settings)
	if result.Name != "Silent Mary" {
		t.Errorf("Expected 'Silent Mary', but got %s", result.Name)
	}
}

// Supports mixed types in the settings list
func TestApplyMixedWithMixedTypes(t *testing.T) {
	// Hoist the sails! Mixed types be no match for us!
	type Ship struct {
		Name  string
		Speed int
	}
	s := &Ship{}
	settings := []interface{}{
		Func[Ship](func(s *Ship) { s.Name = "Queen Anne's Revenge" }),
		func(s *Ship) { s.Speed = 20 },
		Func[Ship](func(s *Ship) { s.Speed += 5 }),
	}
	ApplyMixed(s, settings)
	if s.Name != "Queen Anne's Revenge" || s.Speed != 25 {
		t.Errorf("Expected 'Queen Anne's Revenge' with speed 25, but got %s with speed %d", s.Name, s.Speed)
	}
}
