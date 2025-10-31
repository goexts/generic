package configure_test

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/configure"
)

// Using the same Ship struct from apply_test.go

// TestBuilderWithDefaultValues tests the NewBuilder function with default values
func TestBuilderWithDefaultValues(t *testing.T) {
	var nilShip *Ship // Simple nil pointer for testing

	t.Run("NewBuilder with default value", func(t *testing.T) {
		defaultShip := &Ship{
			Name:  "Default",
			Crew:  10,
			Speed: 5,
		}

		builder := configure.NewBuilder(defaultShip)
		ship, err := builder.Add(
			func(s *Ship) { s.Name = "Enterprise" },
			func(s *Ship) { s.Crew += 10 },
		).Build()

		assert.NoError(t, err)
		assert.Equal(t, "Enterprise", ship.Name)
		assert.Equal(t, 20, ship.Crew)
		assert.Equal(t, 5, ship.Speed)
	})

	t.Run("NewBuilder with nil value", func(t *testing.T) {
		builder := configure.NewBuilder(nilShip)
		ship, err := builder.Add(
			func(s *Ship) { s.Name = "Voyager" },
			func(s *Ship) { s.Crew = 50 },
		).Build()

		assert.NoError(t, err)
		assert.Equal(t, "Voyager", ship.Name)
		assert.Equal(t, 50, ship.Crew)
	})

	t.Run("Type inference with helper function", func(t *testing.T) {
		createBuilder := func(defaultVal *Ship) *configure.Builder[Ship] {
			return configure.NewBuilder(defaultVal)
		}

		builder := createBuilder(nilShip).Add(
			func(s *Ship) { s.Name = "Discovery" },
		)

		ship, err := builder.Build()
		assert.NoError(t, err)
		assert.Equal(t, "Discovery", ship.Name)
	})
}

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

		// Test Apply method (replaces ApplyTo)
		targetShip := &Ship{}
		err := builder.Apply(targetShip)
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
			Add(func(_ *Ship) error { return testErr })

		// Test Apply method with error (replaces ApplyTo)
		err := builder.Apply(&Ship{})
		assert.Error(t, err)

		// Test Build method with error
		_, err = builder.Build()
		assert.Error(t, err)
	})

	t.Run("builder can be used as an option itself", func(t *testing.T) {
		// Create a builder with some options
		innerBuilder := configure.NewBuilder[Ship]().
			Add(func(s *Ship) { s.Name = "Inner" })

		// Use this builder as an option within the NewAny function
		ship, err := configure.NewAny[Ship](
			func(s *Ship) { s.Speed = 99 },
			innerBuilder, // The builder implements ApplierE and can be passed directly
		)

		assert.NoError(t, err)
		assert.Equal(t, "Inner", ship.Name)
		assert.Equal(t, 99, ship.Speed)
	})

	t.Run("builder with pointer type should panic", func(t *testing.T) {
		// Using a pointer type for the builder is not allowed and should panic at NewBuilder.
		assert.Panics(t, func() {
			configure.NewBuilder[*Ship]()
		}, "NewBuilder should panic when C is a pointer type")
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

		// Parameter order changed: factory, builder
		engine, err := configure.Compile(factory, builder)

		assert.NoError(t, err)
		assert.NotNil(t, engine)
		assert.Equal(t, "500 HP", engine.PowerOutput)
	})

	t.Run("fail compilation when config build fails", func(t *testing.T) {
		testErr := errors.New("config build failed")
		builder := configure.NewBuilder[EngineConfig]().
			Add(func(_ *EngineConfig) error { return testErr })

		// Parameter order changed: factory, builder
		_, err := configure.Compile(factory, builder)
		assert.Error(t, err)
	})

	t.Run("fail compilation when factory fails", func(t *testing.T) {
		builder := configure.NewBuilder[EngineConfig]().
			Add(func(c *EngineConfig) { c.Horsepower = 100 }) // Fuel is missing

		// Parameter order changed: factory, builder
		_, err := configure.Compile(factory, builder)
		assert.EqualError(t, err, "fuel type cannot be empty")
	})
}

// TestChainFunctions tests the Chain and ChainE functions from options.go
func TestChainFunctions(t *testing.T) {
	// Define custom option types for testing with Chain/ChainE
	type MyShipOption func(*Ship)
	type MyShipOptionE func(*Ship) error

	t.Run("Chain combines non-error options with type inference", func(t *testing.T) {
		ship := &Ship{}

		// Test type inference with custom option type
		opt1 := MyShipOption(func(s *Ship) { s.Name = "Voyager" })
		opt2 := MyShipOption(func(s *Ship) { s.Crew = 100 })
		opt3 := MyShipOption(func(s *Ship) { s.Speed = 20 })

		// Rely on type inference
		chainedOpt := configure.Chain(opt1, opt2, opt3)

		// Apply the chained option directly
		chainedOpt(ship)

		assert.Equal(t, "Voyager", ship.Name)
		assert.Equal(t, 100, ship.Crew)
		assert.Equal(t, 20, ship.Speed)
	})

	t.Run("Chain works with direct function literals", func(t *testing.T) {
		ship := &Ship{}

		// This should work with type inference
		chainedOpt := configure.Chain(
			func(s *Ship) { s.Name = "Enterprise" },
			func(s *Ship) { s.Crew = 200 },
		)

		chainedOpt(ship)

		assert.Equal(t, "Enterprise", ship.Name)
		assert.Equal(t, 200, ship.Crew)
	})

	t.Run("ChainE combines error-returning options and stops on error", func(t *testing.T) {
		ship := &Ship{}
		expectedErr := errors.New("engine failure")

		// Using custom option type to test type inference
		optE1 := MyShipOptionE(func(s *Ship) error { s.Name = "Enterprise"; return nil })
		optE2 := MyShipOptionE(func(s *Ship) error { s.Status = "Damaged"; return expectedErr })
		optE3 := MyShipOptionE(func(s *Ship) error { s.Speed = 5; return nil }) // Should not be applied

		// Rely on type inference
		chainedOptE := configure.ChainE(optE1, optE2, optE3)
		err := chainedOptE(ship)

		// Check if the error is a ConfigError and contains our expected error
		var configErr *configure.ConfigError
		assert.ErrorAs(t, err, &configErr)
		assert.ErrorIs(t, configErr.Err, expectedErr)

		assert.Equal(t, "Enterprise", ship.Name) // First option applied
		assert.Equal(t, "Damaged", ship.Status)  // Second option applied before error
	})

	t.Run("ChainE combines error-returning options successfully", func(t *testing.T) {
		ship := &Ship{}

		// Create a chain of options that all succeed
		chainedOptE := configure.ChainE(
			func(s *Ship) error { s.Name = "Discovery"; return nil },
			func(s *Ship) error { s.Crew = 200; return nil },
		)

		// Apply the chained option directly
		err := chainedOptE(ship)

		assert.NoError(t, err)
		assert.Equal(t, "Discovery", ship.Name)
		assert.Equal(t, 200, ship.Crew)
	})

	t.Run("Chain works with mixed concrete types", func(t *testing.T) {
		ship := &Ship{}

		// Mixing different but compatible function types
		opt1 := func(s *Ship) { s.Name = "Voyager" }
		var opt2 MyShipOption = func(s *Ship) { s.Crew = 150 }

		// This should work with type inference
		chainedOpt := configure.Chain(opt1, opt2)
		chainedOpt(ship)

		assert.Equal(t, "Voyager", ship.Name)
		assert.Equal(t, 150, ship.Crew)
	})
}
