package configure_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/configure"
)

type Ship struct {
	Name  string
	Speed int
	Crew  int
}

// TestApplyAny is the core test for the reflection-based implementation.
func TestApplyAny(t *testing.T) {
	t.Run("applies mixed options without error", func(t *testing.T) {
		ship := &Ship{}
		opts := []any{
			configure.Option[Ship](func(s *Ship) { s.Name = "Queen Anne's Revenge" }),
			func(s *Ship) error { s.Speed = 20; return nil },
		}
		_, err := configure.ApplyAny(ship, opts)
		assert.NoError(t, err)
		assert.Equal(t, "Queen Anne's Revenge", ship.Name)
		assert.Equal(t, 20, ship.Speed)
	})

	t.Run("stops on first error in mixed options", func(t *testing.T) {
		ship := &Ship{}
		testErr := errors.New("mutiny")
		opts := []any{
			func(s *Ship) { s.Name = "Bounty" },
			func(s *Ship) error { return testErr },
			func(s *Ship) { s.Speed = 10 },
		}
		_, err := configure.ApplyAny(ship, opts)
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, "Bounty", ship.Name)
		assert.Equal(t, 0, ship.Speed)
	})

	// TODO: Test custom defined option types.
	t.Run("handles custom defined option types", func(t *testing.T) {
		type CustomOption func(*Ship)
		withPort := func(p int) CustomOption {
			return func(s *Ship) { s.Crew = p }
		}

		ship := &Ship{}
		opts := []any{withPort(80)}
		_, err := configure.ApplyAny(ship, opts)
		assert.NoError(t, err)
		assert.Equal(t, 80, ship.Crew)
	})

	t.Run("handles alias option types", func(t *testing.T) {
		type AliasOption = func(*Ship)
		withCrew := func(c int) AliasOption {
			return func(s *Ship) { s.Crew = c }
		}

		ship := &Ship{}
		opts := []any{withCrew(120)}
		_, err := configure.ApplyAny(ship, opts)
		assert.NoError(t, err)
		assert.Equal(t, 120, ship.Crew)
	})

	t.Run("handles new type from configure.Option", func(t *testing.T) {
		type ShipOption configure.Option[Ship]
		withName := func(name string) ShipOption {
			return func(s *Ship) {
				s.Name = name
			}
		}

		ship := &Ship{}
		opts := []any{withName("The Flying Dutchman")}
		_, err := configure.ApplyAny(ship, opts)
		assert.NoError(t, err)
		assert.Equal(t, "The Flying Dutchman", ship.Name)
	})

	t.Run("handles new type from configure.OptionE", func(t *testing.T) {
		type ShipOptionE configure.OptionE[Ship]
		withCrewE := func(c int) ShipOptionE {
			return func(s *Ship) error {
				s.Crew = c
				return nil
			}
		}

		ship := &Ship{}
		opts := []any{withCrewE(250)}
		_, err := configure.ApplyAny(ship, opts)
		assert.NoError(t, err)
		assert.Equal(t, 250, ship.Crew)
	})

	t.Run("handles error from new type from configure.OptionE", func(t *testing.T) {
		type ShipOptionE configure.OptionE[Ship]
		testErr := errors.New("iceberg")
		withCrewE := func(c int) ShipOptionE {
			return func(s *Ship) error {
				s.Crew = c
				return testErr
			}
		}

		ship := &Ship{}
		opts := []any{withCrewE(300)}
		_, err := configure.ApplyAny(ship, opts)
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, 300, ship.Crew) // Value should be set before error is returned
	})

	t.Run("returns error for invalid option type", func(t *testing.T) {
		ship := &Ship{}
		opts := []any{"not a function"}
		_, err := configure.ApplyAny(ship, opts)
		assert.Error(t, err)
	})
}

// TestApplyAndApplyE test the type-safe wrappers.
func TestApplyAndApplyE(t *testing.T) {
	t.Run("Apply with standard options", func(t *testing.T) {
		ship := &Ship{}
		configure.Apply(ship, []configure.Option[Ship]{func(s *Ship) { s.Name = "Endeavour" }})
		assert.Equal(t, "Endeavour", ship.Name)
	})

	t.Run("Apply with custom defined option type", func(t *testing.T) {
		type CustomShipOption func(*Ship)
		customOpts := []CustomShipOption{
			func(s *Ship) { s.Speed = 15 },
		}
		ship := &Ship{}
		configure.Apply(ship, customOpts)
		assert.Equal(t, 15, ship.Speed)
	})

	t.Run("ApplyE with successful options", func(t *testing.T) {
		ship := &Ship{}
		_, err := configure.ApplyE(ship, []configure.OptionE[Ship]{func(s *Ship) error { s.Crew = 200; return nil }})
		assert.NoError(t, err)
		assert.Equal(t, 200, ship.Crew)
	})

	t.Run("ApplyE with failing option", func(t *testing.T) {
		testErr := errors.New("cannon exploded")
		ship := &Ship{}
		_, err := configure.ApplyE(ship, []configure.OptionE[Ship]{func(s *Ship) error { return testErr }})
		assert.ErrorIs(t, err, testErr)
	})
}

func TestOptionSet(t *testing.T) {
	t.Run("OptionSet applies all options", func(t *testing.T) {
		ship := &Ship{}
		options := []configure.Option[Ship]{
			func(s *Ship) { s.Name = "Black Pearl" },
			func(s *Ship) { s.Speed = 50 },
		}
		set := configure.OptionSet(options...)
		configure.Apply(ship, []configure.Option[Ship]{set})
		assert.Equal(t, "Black Pearl", ship.Name)
		assert.Equal(t, 50, ship.Speed)
	})

	t.Run("OptionSetE applies all options successfully", func(t *testing.T) {
		ship := &Ship{}
		options := []configure.OptionE[Ship]{
			func(s *Ship) error { s.Name = "Interceptor"; return nil },
			func(s *Ship) error { s.Crew = 10; return nil },
		}
		set := configure.OptionSetE(options...)
		_, err := configure.ApplyE(ship, []configure.OptionE[Ship]{set})
		assert.NoError(t, err)
		assert.Equal(t, "Interceptor", ship.Name)
		assert.Equal(t, 10, ship.Crew)
	})

	t.Run("OptionSetE stops on first error", func(t *testing.T) {
		ship := &Ship{}
		testErr := errors.New("kraken")
		options := []configure.OptionE[Ship]{
			func(s *Ship) error { s.Name = "The Dauntless"; return nil },
			func(s *Ship) error { return testErr },
			func(s *Ship) error { s.Speed = 30; return nil },
		}
		set := configure.OptionSetE(options...)
		_, err := configure.ApplyE(ship, []configure.OptionE[Ship]{set})
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, "The Dauntless", ship.Name)
		assert.Equal(t, 0, ship.Speed)
	})
}
