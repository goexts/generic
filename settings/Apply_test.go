/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package settings

import (
	"testing"
)

// Apply function modifies the input struct with provided settings
func TestApplyModifiesInputStruct(t *testing.T) {
	// Arrr, let's set sail with some initial values!
	type Ship struct {
		Name string
	}
	modifyName := func(s *Ship) {
		s.Name = "Black Pearl"
	}
	ship := &Ship{Name: "Queen Anne's Revenge"}
	ApplyOr(ship, modifyName)
	if ship.Name != "Black Pearl" {
		t.Errorf("Expected ship name to be 'Black Pearl', but got %s", ship.Name)
	}
}

// Apply returns the modified struct when input is not nil
func TestApplyReturnsModifiedStruct(t *testing.T) {
	// Avast! We be testing the return value!
	type Ship struct {
		Name string
	}
	modifyName := []func(*Ship){
		func(s *Ship) {
			s.Name = "Flying Dutchman"
		},
	}
	ship := &Ship{Name: "Jolly Roger"}
	result := Apply(ship, modifyName)
	if result.Name != "Flying Dutchman" {
		t.Errorf("Expected returned ship name to be 'Flying Dutchman', but got %s", result.Name)
	}
}

// Apply processes all settings in the slice
func TestApplyProcessesAllSettings(t *testing.T) {
	// Shiver me timbers! Let's see if all settings be applied!
	type Ship struct {
		Name  string
		Speed int
	}
	modifyName := ApplyFunc[Ship](func(s *Ship) {
		s.Name = "HMS Victory"
	})
	increaseSpeed := ApplyFunc[Ship](func(s *Ship) {
		s.Speed += 10
	})
	ship := &Ship{Name: "Endeavour", Speed: 20}
	Apply(ship, []func(*Ship){modifyName, increaseSpeed})
	if ship.Name != "HMS Victory" || ship.Speed != 30 {
		t.Errorf("Expected ship name 'HMS Victory' and speed 30, but got %s and %d", ship.Name, ship.Speed)
	}
}

// Apply returns nil when input struct is nil
func TestApplyReturnsNilForNilInput(t *testing.T) {
	// Aye, a ghost ship! Let's see if it returns nil!
	type Ship struct {
		Name string
	}
	result := Apply[Ship](nil, nil)
	if result != nil {
		t.Error("Expected nil for nil input, but got a non-nil result")
	}
}

// Apply handles an empty settings slice without errors
func TestApplyHandlesEmptySettingsSlice(t *testing.T) {
	// Arrr, no settings? No problem!
	type Ship struct {
		Name string
	}
	ship := &Ship{Name: "Cutty Sark"}
	result := Apply(ship, []func(*Ship){})
	if result.Name != "Cutty Sark" {
		t.Errorf("Expected ship name 'Cutty Sark', but got %s", result.Name)
	}
}

// Apply correctly processes settings when only one setting is provided
func TestApplySingleSetting(t *testing.T) {
	// One setting to rule them all!
	type Ship struct {
		Name string
	}
	modifyName := func(s *Ship) {
		s.Name = "Santa Maria"
	}
	ship := &Ship{Name: "Pinta"}
	Apply(ship, []func(s *Ship){modifyName})
	if ship.Name != "Santa Maria" {
		t.Errorf("Expected ship name 'Santa Maria', but got %s", ship.Name)
	}
}

// Apply maintains struct integrity when no settings are applied
func TestApplyMaintainsStructIntegrity(t *testing.T) {
	// No changes? No worries, matey!
	type Ship struct {
		Name string
		Crew int
	}
	ship := &Ship{Name: "Mayflower", Crew: 102}
	Apply(ship, []func(s *Ship){})
	if ship.Name != "Mayflower" || ship.Crew != 102 {
		t.Errorf("Expected ship name 'Mayflower' and crew 102, but got %s and %d", ship.Name, ship.Crew)
	}
}
