package configure_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/configure"
)

// Ship is a simple struct used for basic tests.

type Ship struct {
	Name  string
	Speed int
	Crew  int
}

// MegaShip is a large struct used in TestApplyAny_Variations to ensure each option
// type can modify a unique field, making verification straightforward.
type MegaShip struct {
	F1RawFunc         string
	F2RawFuncE        string
	F3Option          string
	F4OptionE         string
	F5Applier         string
	F6ApplierE        string
	F7CustomFunc      string
	F8CustomFuncE     string
	F9CustomOption    string
	F10CustomOptionE  string
	F11CustomApplier  string
	F12CustomApplierE string
	F13AliasFunc      string
	F14AliasFuncE     string
	F15AliasOption    string
	F16AliasOptionE   string
	F17AliasApplier   string
	F18AliasApplierE  string
}

// --- Helper types for TestApplyAny_Variations ---

type ( // Base Applier Implementations
	BaseApplierImpl  struct{}
	BaseApplierEImpl struct{}
)

func (a BaseApplierImpl) Apply(s *MegaShip)        { s.F5Applier = "ok" }
func (a BaseApplierEImpl) Apply(s *MegaShip) error { s.F6ApplierE = "ok"; return nil }

type ( // Custom Applier Implementations
	CustomApplierImpl  struct{}
	CustomApplierEImpl struct{}
)

func (a CustomApplierImpl) Apply(s *MegaShip)        { s.F11CustomApplier = "ok" }
func (a CustomApplierEImpl) Apply(s *MegaShip) error { s.F12CustomApplierE = "ok"; return nil }

type ( // Alias Applier Implementations
	AliasApplierImpl  struct{}
	AliasApplierEImpl struct{}
)

func (a AliasApplierImpl) Apply(s *MegaShip)        { s.F17AliasApplier = "ok" }
func (a AliasApplierEImpl) Apply(s *MegaShip) error { s.F18AliasApplierE = "ok"; return nil }

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
			func(_ *Ship) error { return testErr },
			func(s *Ship) { s.Speed = 10 },
		}
		_, err := configure.ApplyAny(ship, opts)
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, "Bounty", ship.Name)
		assert.Equal(t, 0, ship.Speed)
	})

	t.Run("returns error for invalid option type", func(t *testing.T) {
		ship := &Ship{}
		opts := []any{"not a function"}
		_, err := configure.ApplyAny(ship, opts)
		assert.Error(t, err)
	})
}

// TestApplyAny_Variations provides the most comprehensive testing for ApplyAny,
// ensuring that all supported option types, including their custom and alias variations,
// are handled correctly.
func TestApplyAny_Variations(t *testing.T) {
	// --- 1. Base Types ---
	opt1 := func(s *MegaShip) { s.F1RawFunc = "ok" }
	opt2 := func(s *MegaShip) error { s.F2RawFuncE = "ok"; return nil }
	opt3 := configure.Option[MegaShip](func(s *MegaShip) { s.F3Option = "ok" })
	opt4 := configure.OptionE[MegaShip](func(s *MegaShip) error { s.F4OptionE = "ok"; return nil })
	opt5 := BaseApplierImpl{}
	opt6 := BaseApplierEImpl{}

	// --- 2. Custom Types (type MyType ...)
	type CustomFunc func(*MegaShip)
	opt7 := CustomFunc(func(s *MegaShip) { s.F7CustomFunc = "ok" })
	type CustomFuncE func(*MegaShip) error
	opt8 := CustomFuncE(func(s *MegaShip) error { s.F8CustomFuncE = "ok"; return nil })
	type CustomOption configure.Option[MegaShip]
	opt9 := CustomOption(func(s *MegaShip) { s.F9CustomOption = "ok" })
	type CustomOptionE configure.OptionE[MegaShip]
	opt10 := CustomOptionE(func(s *MegaShip) error { s.F10CustomOptionE = "ok"; return nil })
	opt11 := CustomApplierImpl{}
	opt12 := CustomApplierEImpl{}

	// --- 3. Type Aliases (type MyAlias = ...)
	opt13 := func(s *MegaShip) { s.F13AliasFunc = "ok" }
	opt14 := func(s *MegaShip) error { s.F14AliasFuncE = "ok"; return nil }
	type AliasOption = configure.Option[MegaShip]
	opt15 := AliasOption(func(s *MegaShip) { s.F15AliasOption = "ok" })
	type AliasOptionE = configure.OptionE[MegaShip]
	opt16 := AliasOptionE(func(s *MegaShip) error { s.F16AliasOptionE = "ok"; return nil })
	type AliasApplier = AliasApplierImpl
	opt17 := AliasApplier{}
	type AliasApplierE = AliasApplierEImpl
	opt18 := AliasApplierE{}

	t.Run("successfully applies a mix of all 18 variations", func(t *testing.T) {
		ship := &MegaShip{}
		opts := []any{
			opt1, opt2, opt3, opt4, opt5, opt6,
			opt7, opt8, opt9, opt10, opt11, opt12,
			opt13, opt14, opt15, opt16, opt17, opt18,
		}

		_, err := configure.ApplyAny(ship, opts)

		assert.NoError(t, err)
		assert.Equal(t, "ok", ship.F1RawFunc)
		assert.Equal(t, "ok", ship.F2RawFuncE)
		assert.Equal(t, "ok", ship.F3Option)
		assert.Equal(t, "ok", ship.F4OptionE)
		assert.Equal(t, "ok", ship.F5Applier)
		assert.Equal(t, "ok", ship.F6ApplierE)
		assert.Equal(t, "ok", ship.F7CustomFunc)
		assert.Equal(t, "ok", ship.F8CustomFuncE)
		assert.Equal(t, "ok", ship.F9CustomOption)
		assert.Equal(t, "ok", ship.F10CustomOptionE)
		assert.Equal(t, "ok", ship.F11CustomApplier)
		assert.Equal(t, "ok", ship.F12CustomApplierE)
		assert.Equal(t, "ok", ship.F13AliasFunc)
		assert.Equal(t, "ok", ship.F14AliasFuncE)
		assert.Equal(t, "ok", ship.F15AliasOption)
		assert.Equal(t, "ok", ship.F16AliasOptionE)
		assert.Equal(t, "ok", ship.F17AliasApplier)
		assert.Equal(t, "ok", ship.F18AliasApplierE)
	})

	t.Run("stops on first error in a comprehensive mix", func(t *testing.T) {
		ship := &MegaShip{}
		testErr := errors.New("engine failure")
		failingOpt := configure.OptionE[MegaShip](func(_ *MegaShip) error { return testErr })

		opts := []any{
			opt1,       // Should apply
			failingOpt, // Should fail here
			opt3,       // Should NOT apply
		}

		_, err := configure.ApplyAny(ship, opts)

		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, "ok", ship.F1RawFunc, "Should be applied before error")
		assert.Empty(t, ship.F3Option, "Should not be applied after error")
	})
}

// TestApplyAndApplyE test the type-safe wrappers.
func TestApplyAndApplyE(t *testing.T) {
	t.Run("Apply with standard options", func(t *testing.T) {
		ship := &Ship{}
		configure.Apply(ship, []configure.Option[Ship]{func(s *Ship) { s.Name = "Endeavor" }})
		assert.Equal(t, "Endeavor", ship.Name)
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
		_, err := configure.ApplyE(ship, []configure.OptionE[Ship]{func(_ *Ship) error { return testErr }})
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
			func(_ *Ship) error { return testErr },
			func(s *Ship) error { s.Speed = 30; return nil },
		}
		set := configure.OptionSetE(options...)
		_, err := configure.ApplyE(ship, []configure.OptionE[Ship]{set})
		assert.ErrorIs(t, err, testErr)
		assert.Equal(t, "The Dauntless", ship.Name)
		assert.Equal(t, 0, ship.Speed)
	})
}

func TestNewConstructors(t *testing.T) {
	t.Run("NewWith creates object with standard options", func(t *testing.T) {
		ship := configure.NewWith[Ship](
			func(s *Ship) { s.Name = "Millennium Falcon" },
			func(s *Ship) { s.Speed = 99 },
		)
		assert.NotNil(t, ship)
		assert.Equal(t, "Millennium Falcon", ship.Name)
		assert.Equal(t, 99, ship.Speed)
	})

	t.Run("New (non-variadic) creates object with standard options", func(t *testing.T) {
		opts := []configure.Option[Ship]{
			func(s *Ship) { s.Name = "Normandy" },
		}
		ship := configure.New(opts)
		assert.Equal(t, "Normandy", ship.Name)
	})

	t.Run("NewWith handles empty options", func(t *testing.T) {
		ship := configure.NewWith[Ship]()
		assert.NotNil(t, ship)
		assert.Equal(t, "", ship.Name) // Should be zero-valued
	})

	t.Run("NewWithE creates object with successful error-returning options", func(t *testing.T) {
		ship, err := configure.NewWithE[Ship](
			func(s *Ship) error { s.Name = "Endeavour"; return nil },
		)
		assert.NoError(t, err)
		assert.Equal(t, "Endeavour", ship.Name)
	})

	t.Run("NewE (non-variadic) handles failing option", func(t *testing.T) {
		testErr := errors.New("construction failed")
		opts := []configure.OptionE[Ship]{
			func(s *Ship) error { s.Name = "Titanic"; return nil },
			func(_ *Ship) error { return testErr },
		}
		ship, err := configure.NewE(opts)
		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, ship) // Object should not be returned on error
	})

	t.Run("NewWithE handles failing option", func(t *testing.T) {
		testErr := errors.New("construction failed")
		ship, err := configure.NewWithE[Ship](
			func(s *Ship) error { s.Name = "Titanic"; return nil },
			func(_ *Ship) error { return testErr },
		)
		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, ship) // Object should not be returned on error
	})

	t.Run("NewAny creates object with mixed options", func(t *testing.T) {
		ship, err := configure.NewAny[Ship](
			func(s *Ship) { s.Name = "Mixed" },
			func(s *Ship) error { s.Speed = 10; return nil },
		)
		assert.NoError(t, err)
		assert.Equal(t, "Mixed", ship.Name)
		assert.Equal(t, 10, ship.Speed)
	})

	t.Run("NewAny handles failing mixed option", func(t *testing.T) {
		testErr := errors.New("any construction failed")
		ship, err := configure.NewAny[Ship](
			func(s *Ship) { s.Name = "Partial" },
			func(_ *Ship) error { return testErr },
		)
		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, ship) // Object should not be returned on error
	})

	t.Run("NewWith using OptionSet", func(t *testing.T) {
		set := configure.OptionSet(
			func(s *Ship) { s.Name = "SetShip" },
			func(s *Ship) { s.Speed = 42 },
		)
		ship := configure.NewWith(set)
		assert.Equal(t, "SetShip", ship.Name)
		assert.Equal(t, 42, ship.Speed)
	})

	t.Run("NewWithE using OptionSetE with error", func(t *testing.T) {
		testErr := errors.New("set error")
		set := configure.OptionSetE(
			func(s *Ship) error { s.Name = "Should be set"; return nil },
			func(_ *Ship) error { return testErr },
		)
		ship, err := configure.NewWithE(set)
		assert.ErrorIs(t, err, testErr)
		assert.Nil(t, ship)
	})
}
