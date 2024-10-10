package settings

import (
	"testing"
)

// Apply function correctly modifies the input object
func TestApplyModifiesInputObject(t *testing.T) {
	// Arrr! Let's see if this here function can change the course of our ship!
	type Ship struct {
		Name string
	}
	changeName := func(s *Ship) {
		s.Name = "Black Pearl"
	}
	setting := Setting[Ship](changeName)
	ship := &Ship{Name: "Queen Anne's Revenge"}

	setting.Apply(ship)

	if ship.Name != "Black Pearl" {
		t.Errorf("Expected ship name to be 'Black Pearl', but got %s", ship.Name)
	}
}

// Apply function executes without errors for valid input
func TestApplyExecutesWithoutErrors(t *testing.T) {
	// Avast! We be testing if the seas be calm and error-free!
	type Ship struct {
		Name string
	}
	changeName := func(s *Ship) {
		s.Name = "Flying Dutchman"
	}
	setting := Setting[Ship](changeName)
	ship := &Ship{Name: "Jolly Roger"}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Apply panicked with valid input: %v", r)
		}
	}()

	setting.Apply(ship)
}

// Apply function handles multiple settings in sequence
func TestApplyHandlesMultipleSettings(t *testing.T) {
	// Shiver me timbers! Can it handle a fleet of changes?
	type Ship struct {
		Name  string
		Speed int
	}
	changeName := func(s *Ship) {
		s.Name = "HMS Victory"
	}
	increaseSpeed := func(s *Ship) {
		s.Speed += 10
	}

	settings := []Setting[Ship]{changeName, increaseSpeed}
	ship := &Ship{Name: "Endeavour", Speed: 20}

	for _, setting := range settings {
		setting.Apply(ship)
	}

	if ship.Name != "HMS Victory" || ship.Speed != 30 {
		t.Errorf("Expected ship name 'HMS Victory' and speed 30, but got %s and %d", ship.Name, ship.Speed)
	}
}

// Apply function handles nil input without crashing
func TestApplyHandlesNilInput(t *testing.T) {
	// Blimey! Let's see if it sinks when there's no ship!
	type Ship struct {
		Name string
	}
	changeName := func(s *Ship) {
		s.Name = "Ghost Ship"
	}
	setting := Setting[Ship](changeName)

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Apply panicked with nil input: %v", r)
		}
	}()

	setting.Apply(nil)
}

// Apply function processes empty settings list gracefully
func TestApplyProcessesEmptySettingsList(t *testing.T) {
	// Arrr! No changes be made when there be no orders!
	type Ship struct {
		Name string
	}

	var settings []Setting[Ship]
	ship := &Ship{Name: "Silent Mary"}

	for _, setting := range settings {
		setting.Apply(ship)
	}

	if ship.Name != "Silent Mary" {
		t.Errorf("Expected ship name to remain 'Silent Mary', but got %s", ship.Name)
	}
}

// Apply function handles settings that do not modify the object
func TestApplyHandlesNoOpSettings(t *testing.T) {
	// Aye, let's see if it stays the course when no changes be made!
	type Ship struct {
		Name string
	}

	noOp := func(s *Ship) {}

	setting := Setting[Ship](noOp)
	ship := &Ship{Name: "Interceptor"}

	setting.Apply(ship)

	if ship.Name != "Interceptor" {
		t.Errorf("Expected ship name to remain 'Interceptor', but got %s", ship.Name)
	}
}

// Apply function maintains object integrity after application
func TestApplyMaintainsObjectIntegrity(t *testing.T) {
	// Yo ho ho! Make sure the ship stays afloat after changes!
	type Ship struct {
		Name    string
		Cannons int
	}

	addCannons := func(s *Ship) {
		s.Cannons += 2
	}

	setting := Setting[Ship](addCannons)
	ship := &Ship{Name: "Blackbeard's Revenge", Cannons: 10}

	setting.Apply(ship)

	if ship.Cannons != 12 || ship.Name != "Blackbeard's Revenge" {
		t.Errorf("Expected cannons to be 12 and name 'Blackbeard's Revenge', but got %d and %s", ship.Cannons, ship.Name)
	}
}
