package configure_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/configure"
)

// Using the same Ship struct from apply_test.go

func TestWithValidation(t *testing.T) {
	t.Run("validation succeeds", func(t *testing.T) {
		ship := &Ship{Name: "Endurance", Crew: 28}
		validator := func(s *Ship) error {
			if s.Crew > 0 {
				return nil // Success
			}
			return errors.New("ship has no crew")
		}

		validationOpt := configure.WithValidation(validator)
		err := validationOpt(ship)
		assert.NoError(t, err)
	})

	t.Run("validation fails", func(t *testing.T) {
		ship := &Ship{Name: "Ghost Ship", Crew: 0}
		testErr := errors.New("ship has no crew")
		validator := func(s *Ship) error {
			if s.Crew > 0 {
				return nil
			}
			return testErr
		}

		validationOpt := configure.WithValidation(validator)
		err := validationOpt(ship)
		assert.ErrorIs(t, err, testErr)
	})

	t.Run("integration with ApplyE", func(t *testing.T) {
		ship := &Ship{}
		testErr := errors.New("ship name cannot be empty")

		opts := []configure.OptionE[Ship]{
			func(s *Ship) error { s.Crew = 10; return nil },
			configure.WithValidation(func(s *Ship) error {
				if s.Name == "" {
					return testErr
				}
				return nil
			}),
		}

		_, err := configure.ApplyE(ship, opts)
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, 10, ship.Crew) // The option before validation should have been applied
	})
}
