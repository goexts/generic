// Package configure provides utilities for applying functional options to objects.
package configure

import (
	"errors"
	"fmt"
)

type ErrorCode int

// Error types for configuration errors
const (
	ErrUnsupportedType ErrorCode = iota
	ErrExecutionFailed
	ErrEmptyTargetValue
)

// ConfigError is a custom error type for the configure package.
// It provides more context than a standard error.
type ConfigError struct {
	Code       ErrorCode // Error category
	TypeString string    // Setting type info
	Err        error     // Original error
}

func newConfigError(code ErrorCode, setting any, err error) *ConfigError {
	return &ConfigError{
		Code:       code,
		TypeString: fmt.Sprintf("%T", setting),
		Err:        err,
	}
}

// Unwrap allows for unwrapping the original error, making it compatible with errors.Is and errors.As.
func (e *ConfigError) Unwrap() error {
	return e.Err
}

// Error returns the string representation of the ConfigError.
func (e *ConfigError) Error() string {
	switch e.Code {
	case ErrUnsupportedType:
		return fmt.Sprintf("unsupported option type: %s", e.TypeString)
	case ErrExecutionFailed:
		return fmt.Sprintf("option apply failed [type:%s]: %v", e.TypeString, e.Err)
	case ErrEmptyTargetValue:
		return "target for configuration cannot be nil"
	default:
		return "unknown config error"
	}
}

// IsConfigError checks if the given error is a *ConfigError instance.
func IsConfigError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError)
}

// IsUnsupportedTypeError checks if the error is a ConfigError with the code ErrUnsupportedType.
func IsUnsupportedTypeError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrUnsupportedType
}

// IsExecutionFailedError checks if the error is a ConfigError with the code ErrExecutionFailed.
func IsExecutionFailedError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrExecutionFailed
}

// IsEmptyTargetValueError checks if the error is a ConfigError with the code ErrEmptyTargetValue.
func IsEmptyTargetValueError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrEmptyTargetValue
}
