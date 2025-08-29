package configure_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goexts/generic/configure"
)

func TestErrorHandling(t *testing.T) {
	// Using the same Ship struct from other tests

	t.Run("ApplyE returns wrapped execution error", func(t *testing.T) {
		originalErr := errors.New("engine failure")
		failingOpt := func(_ *Ship) error { return originalErr }

		_, err := configure.ApplyE(&Ship{}, []configure.OptionE[Ship]{failingOpt})

		// General error checks
		assert.Error(t, err)
		assert.True(t, configure.IsConfigError(err))
		assert.True(t, configure.IsExecutionFailedError(err))
		assert.False(t, configure.IsUnsupportedTypeError(err))

		// Check if we can unwrap the original error
		assert.ErrorIs(t, err, originalErr)

		// Check the specific error type and its contents
		var configErr *configure.ConfigError
		assert.True(t, errors.As(err, &configErr))
		assert.Equal(t, configure.ErrExecutionFailed, configErr.Code)
		assert.NotZero(t, configErr.TypeString)
	})

	t.Run("ApplyAny returns wrapped execution error", func(t *testing.T) {
		originalErr := errors.New("cannon misfire")
		failingOpt := configure.OptionE[Ship](func(_ *Ship) error { return originalErr })

		_, err := configure.ApplyAnyWith(&Ship{}, failingOpt)

		assert.Error(t, err)
		assert.True(t, configure.IsExecutionFailedError(err))
		assert.ErrorIs(t, err, originalErr)
	})

	t.Run("ApplyAny returns unsupported type error", func(t *testing.T) {
		invalidOpt := "not a function"
		_, err := configure.ApplyAnyWith(&Ship{}, invalidOpt)

		// General error checks
		assert.Error(t, err)
		assert.True(t, configure.IsConfigError(err))
		assert.True(t, configure.IsUnsupportedTypeError(err))
		assert.False(t, configure.IsExecutionFailedError(err))

		// Check the error message
		expectedMsg := fmt.Sprintf("unsupported option type: %T", invalidOpt)
		assert.Equal(t, expectedMsg, err.Error())

		// Check that there is no underlying error
		assert.Nil(t, errors.Unwrap(err))
	})
}
