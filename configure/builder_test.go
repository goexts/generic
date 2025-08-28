package configure_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/goexts/generic/configure"
	"github.com/stretchr/testify/assert"
)

// Using the same Ship struct from apply_test.go

func TestBuilder(t *testing.T) {
	t.Run("build and apply options for the same type", func(t *testing.T) {
		// Create a builder and add various types of options
		builder := configure.NewBuilder[Ship]().
			Add(
				configure.Option[Ship](func(s *Ship) { s.Name = "Argo" }),
				func(s *Ship) { s.Crew = 50 },
			).
			AddWhen(true, func(s *Ship) { s.Speed = 15 }).
			AddWhen(false, func(s *Ship) { s.Name = "Should not be applied" })

		// Test ApplyTo method
		targetShip := &Ship{}
		_, err := builder.ApplyTo(targetShip)
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
		builder := configure.NewBuilder[Ship]().
			Add(func(s *Ship) error { return testErr })

		// Test ApplyTo method with error
		_, err := builder.ApplyTo(&Ship{})
		assert.Error(t, err)

		// Test Build method with error
		_, err = builder.Build()
		assert.Error(t, err)
	})

	t.Run("builder can be used as an option itself", func(t *testing.T) {
		// Create a builder with some options
		innerBuilder := configure.NewBuilder[Ship]().
			Add(func(s *Ship) { s.Name = "Inner" })

		// Use this builder as an option within the New function
		ship, err := configure.New[Ship](
			func(s *Ship) { s.Speed = 99 },
			innerBuilder, // The builder implements ApplierE and can be passed directly
		)

		assert.NoError(t, err)
		assert.Equal(t, "Inner", ship.Name)
		assert.Equal(t, 99, ship.Speed)
	})
}

func TestCompile(t *testing.T) {
	// Define a config struct and a product struct
	type EngineConfig struct {
		Horsepower int
		Fuel       string
	}
	type Engine struct {
		PowerOutput string
	}

	// Define a factory to "compile" the config into a product
	factory := func(c *EngineConfig) (*Engine, error) {
		if c.Fuel == "" {
			return nil, errors.New("fuel type cannot be empty")
		}
		return &Engine{PowerOutput: strconv.Itoa(c.Horsepower) + " HP"}, nil
	}

	t.Run("successfully compile a product from a config", func(t *testing.T) {
		builder := configure.NewBuilder[EngineConfig]().
			Add(func(c *EngineConfig) { c.Horsepower = 500 }).
			Add(func(c *EngineConfig) { c.Fuel = "Gasoline" })

		engine, err := configure.Compile(builder, factory)

		assert.NoError(t, err)
		assert.NotNil(t, engine)
		assert.Equal(t, "500 HP", engine.PowerOutput)
	})

	t.Run("fail compilation when config build fails", func(t *testing.T) {
		testErr := errors.New("config build failed")
		builder := configure.NewBuilder[EngineConfig]().
			Add(func(c *EngineConfig) error { return testErr })

		_, err := configure.Compile(builder, factory)
		assert.Error(t, err)
	})

	t.Run("fail compilation when factory fails", func(t *testing.T) {
		factoryErr := errors.New("fuel type cannot be empty")
		builder := configure.NewBuilder[EngineConfig]().
			Add(func(c *EngineConfig) { c.Horsepower = 100 }) // Fuel is missing

		_, err := configure.Compile(builder, factory)
		assert.ErrorIs(t, err, factoryErr)
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
		assert.Error(t, err)
	})
}
