package settings

import (
	"testing"
)

// ApplyOr applies settings to a non-nil struct
func TestApplyOrAppliesSettingsToNonNilStruct(t *testing.T) {
	// Arrr, let's set sail with a non-nil struct!
	type Ship struct {
		Name string
	}
	s := &Ship{Name: "Black Pearl"}

	setting := func(s *Ship) {
		s.Name = "Flying Dutchman"
	}
	type ShipSetting ApplyFunc[Ship]

	setting2 := ShipSetting(func(s *Ship) {
		s.Name = "Queen Anne's Revenge"
	})

	result := ApplyOr(s, setting, setting2)
	if result.Name != "Queen Anne's Revenge" {
		t.Errorf("Expected 'Queen Anne's Revenge', got %s", result.Name)
	}
}

// ApplyOr returns the modified struct after applying settings
func TestApplyOrReturnsModifiedStruct(t *testing.T) {
	// Aye, the struct be modified, or ye walk the plank!
	type Ship struct {
		Name string
	}
	s := &Ship{Name: "Jolly Roger"}
	setting := func(s *Ship) {
		s.Name = "Queen Anne's Revenge"
	}
	result := ApplyOr(s, setting)
	if result != s {
		t.Error("Expected the same struct reference to be returned")
	}
}

// ApplyOr can handle multiple settings being applied in sequence
func TestApplyOrHandlesMultipleSettings(t *testing.T) {
	// Avast! Multiple settings be comin' aboard!
	type Ship struct {
		Name  string
		Speed int
	}
	s := &Ship{Name: "Interceptor", Speed: 10}
	setting1 := func(s *Ship) {
		s.Name = "Endeavour"
	}
	setting2 := func(s *Ship) {
		s.Speed = 20
	}
	result := ApplyOr(s, setting1, setting2)
	if result.Name != "Endeavour" || result.Speed != 20 {
		t.Errorf("Expected 'Endeavour' and speed 20, got %s and speed %d", result.Name, result.Speed)
	}
}

// ApplyOr handles a nil struct input gracefully
func TestApplyOrHandlesNilStructInput(t *testing.T) {
	// Shiver me timbers! A nil struct be no problem!
	type Ship struct {
		Name string
	}
	var s *Ship = nil
	setting := ApplyFunc[Ship](func(s *Ship) {
		s.Name = "Ghost Ship"
	})
	result := ApplyOr(s, setting)
	if result != nil {
		t.Error("Expected nil when input struct is nil")
	}
}

// ApplyOr processes an empty list of settings without error
func TestApplyOrProcessesEmptySettingsList(t *testing.T) {
	// Arrr, no settings? No worries, matey!
	type Ship struct {
		Name string
	}
	s := &Ship{Name: "Silent Mary"}
	result := ApplyOr(s)
	if result.Name != "Silent Mary" {
		t.Errorf("Expected 'Silent Mary', got %s", result.Name)
	}
}

// ApplyOr handles settings that do not modify the struct
func TestApplyOrHandlesNoOpSettings(t *testing.T) {
	// Aye, some settings be as useful as a landlubber!
	type Ship struct {
		Name string
	}
	s := &Ship{Name: "Davy Jones"}
	noopSetting := func(s *Ship) {}
	result := ApplyOr(s, noopSetting)
	if result.Name != "Davy Jones" {
		t.Errorf("Expected 'Davy Jones', got %s", result.Name)
	}
}

// ApplyOr should maintain the original struct if no settings are applied
func TestApplyOrMaintainsOriginalStructIfNoSettingsApplied(t *testing.T) {
	// Blimey! No changes should be made to the original vessel!
	type Ship struct {
		Name string
		Crew int
	}
	s := &Ship{Name: "Adventure Galley", Crew: 100}
	result := ApplyOr(s)
	if result.Name != "Adventure Galley" || result.Crew != 100 {
		t.Errorf("Expected 'Adventure Galley' with crew 100, got %s with crew %d", result.Name, result.Crew)
	}
}
