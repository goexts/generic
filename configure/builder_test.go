package configure_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/configure"
)

// Using the same Ship struct from apply_test.go

func TestBuilder(t *testing.T) {
	t.Run("build and apply options", func(t *testing.T) {
		// Create a builder and add various types of options
		builder := configure.NewBuilder[Ship]()
		builder.With(
			configure.Option[Ship](func(s *Ship) { s.Name = "Argo" }),
			func(s *Ship) { s.Crew = 50 },
		)
		builder.WithWhen(true, func(s *Ship) { s.Speed = 15 })
		builder.WithWhen(false, func(s *Ship) { s.Name = "Should not be applied" })

		// Test ApplyTo method
		targetShip := &Ship{}
		err := builder.ApplyTo(targetShip)
		assert.NoError(t, err)
		assert.Equal(t, "Argo", targetShip.Name)
		assert.Equal(t, 50, targetShip.Crew)
		assert.Equal(t, 15, targetShip.Speed)

		// Test Build method
		builtShip, err := builder.Build()
		assert.NoError(t, err)
		assert.Equal(t, "Argo", builtShip.Name)
		assert.Equal(t, 50, builtShip.Crew)
		assert.Equal(t, 15, builtShip.Speed)
	})

	t.Run("builder with failing option", func(t *testing.T) {
		testErr := errors.New("cannot build")
		builder := configure.NewBuilder[Ship]()
		builder.With(func(s *Ship) error { return testErr })

		// Test ApplyTo method with error
		err := builder.ApplyTo(&Ship{})
		assert.ErrorIs(t, err, testErr)

		// Test Build method with error
		_, err = builder.Build()
		assert.ErrorIs(t, err, testErr)
	})
}

func TestNew(t *testing.T) {
	t.Run("create new object with options", func(t *testing.T) {
		ship, err := configure.New[Ship](
			func(s *Ship) { s.Name = "Millennium Falcon" },
			func(s *Ship) { s.Speed = 99 },
		)
		assert.NoError(t, err)
		assert.NotNil(t, ship)
		assert.Equal(t, "Millennium Falcon", ship.Name)
		assert.Equal(t, 99, ship.Speed)
	})

	t.Run("create new object with failing option", func(t *testing.T) {
		testErr := errors.New("construction failed")
		_, err := configure.New[Ship](func(s *Ship) error { return testErr })
		assert.ErrorIs(t, err, testErr)
	})
}
