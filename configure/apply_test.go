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
