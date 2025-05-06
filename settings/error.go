// Package settings implements the functions, types, and interfaces for the module.
package settings

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

type ConfigError struct {
	Code       ErrorCode // Error category
	TypeString string    // Setting type info
	Err        error     // Original error
}

func newConfigError(t ErrorCode, setting any, err error) *ConfigError {
	return &ConfigError{
		Code:       t,
		TypeString: fmt.Sprintf("%T", setting),
		Err:        err,
	}
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

// Error message display is enhanced
func (e *ConfigError) Error() string {
	switch e.Code {
	case ErrUnsupportedType:
		return fmt.Sprintf("unsupported config type: %s", e.TypeString)
	case ErrExecutionFailed:
		return fmt.Sprintf("config apply failed [type:%s]: %v", e.TypeString, e.Err)
	case ErrEmptyTargetValue:
		return "target value is empty"
	default:
		return "unknown config error"
	}
}

func wrapIfNeeded(err error, setting any) error {
	if err == nil {
		return nil
	}

	var ce *ConfigError
	if errors.As(err, &ce) {
		return err
	}

	return newConfigError(ErrExecutionFailed, setting, err)
}

// IsConfigError checks if the given error is a *ConfigError instance.
// Parameters:
//   - err: error to be checked
//
// Returns:
//   - true if err is *ConfigError type, false otherwise
func IsConfigError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError)
}

func IsUnsupportedTypeError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrUnsupportedType
}

func IsEmptyTargetValueError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrEmptyTargetValue
}

func IsExecutionFailedError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrExecutionFailed
}
