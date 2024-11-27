/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package settings

import (
	"testing"
)

// Applies a list of ApplyFunc or func to a struct
func TestApplyAnyWithApplyFuncAndFunc(t *testing.T) {
	// Arrr, let's set sail with some settings!
	type Ship struct {
		Name string
	}
	s := &Ship{}
	settings := []interface{}{
		ApplyFunc[Ship](func(s *Ship) { s.Name = "Black Pearl" }),
		func(s *Ship) { s.Name += " - The Fastest" },
	}
	ApplyAny(s, settings)
	if s.Name != "Black Pearl - The Fastest" {
		t.Errorf("Expected 'Black Pearl - The Fastest', but got %s", s.Name)
	}
}

// Returns the modified struct after applying settings
func TestApplyAnyReturnsModifiedStruct(t *testing.T) {
	// Aye, the ship be modified after the storm!
	type Ship struct {
		Speed int
	}
	s := &Ship{Speed: 10}
	settings := []interface{}{
		func(s *Ship) { s.Speed += 5 },
	}
	result := ApplyAny(s, settings)
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

// Handles ApplySetting interface implementations correctly
func TestApplyAnyWithApplySettingInterface(t *testing.T) {
	s := &Ship{}
	settings := []interface{}{
		CrewSetting{},
	}
	ApplyAny(s, settings)
	if s.Crew != 100 {
		t.Errorf("Expected crew 100, but got %d", s.Crew)
	}
}

// Returns nil if the input struct pointer is nil
func TestApplyAnyWithNilStruct(t *testing.T) {
	// Shiver me timbers! A nil ship be no ship at all!
	type Ship struct {
		Name string
	}
	var s *Ship = nil
	settings := []interface{}{
		func(s *Ship) { s.Name = "Flying Dutchman" },
	}
	result := ApplyAny(s, settings)
	if result != nil {
		t.Error("Expected nil, but got a non-nil result")
	}
}

// Panics when an invalid setting type is provided
func TestApplyAnyPanicsOnInvalidType(t *testing.T) {
	// Arrr! Beware the kraken of invalid types!
	type Ship struct {
		Name string
	}
	s := &Ship{}
	settings := []interface{}{
		"Invalid Type",
	}
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic, but did not panic")
		}
	}()
	ApplyAny(s, settings)
}

// Handles an empty list of settings without error
func TestApplyAnyWithEmptySettingsList(t *testing.T) {
	// Aye, an empty list be as harmless as a calm sea!
	type Ship struct {
		Name string
	}
	s := &Ship{Name: "Silent Mary"}
	settings := []interface{}{}
	result := ApplyAny(s, settings)
	if result.Name != "Silent Mary" {
		t.Errorf("Expected 'Silent Mary', but got %s", result.Name)
	}
}

// Supports mixed types in the settings list
func TestApplyAnyWithMixedTypes(t *testing.T) {
	// Hoist the sails! Mixed types be no match for us!
	type Ship struct {
		Name  string
		Speed int
	}
	s := &Ship{}
	settings := []interface{}{
		ApplyFunc[Ship](func(s *Ship) { s.Name = "Queen Anne's Revenge" }),
		func(s *Ship) { s.Speed = 20 },
		ApplyFunc[Ship](func(s *Ship) { s.Speed += 5 }),
	}
	ApplyAny(s, settings)
	if s.Name != "Queen Anne's Revenge" || s.Speed != 25 {
		t.Errorf("Expected 'Queen Anne's Revenge' with speed 25, but got %s with speed %d", s.Name, s.Speed)
	}
}
