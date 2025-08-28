package configure

import (
	"errors"
	"fmt"
)

// ErrorCode defines the specific category of a configuration error.
type ErrorCode int

// Error codes for specific configuration failures.
const (
	// ErrUnsupportedType indicates that an option's type is not supported by
	// the ApplyAny function.
	ErrUnsupportedType ErrorCode = iota

	// ErrExecutionFailed indicates that an option function returned an error
	// during its execution.
	ErrExecutionFailed

	// ErrEmptyTargetValue indicates that a nil pointer was passed as the target
	// for configuration.
	ErrEmptyTargetValue
)

// ConfigError is a custom error type for the configure package.
// It wraps an original error while providing additional context, such as the
// type of option that caused the failure and a specific error code.
type ConfigError struct {
	// Code is the category of the error.
	Code ErrorCode
	// TypeString is the string representation of the option's type.
	TypeString string
	// Err is the underlying error, if any.
	Err error
}

// newConfigError creates a new ConfigError.
func newConfigError(code ErrorCode, setting any, err error) *ConfigError {
	return &ConfigError{
		Code:       code,
		TypeString: fmt.Sprintf("%T", setting),
		Err:        err,
	}
}

// Unwrap makes ConfigError compatible with the standard library's errors.Is
// and errors.As functions, allowing for proper error chain inspection.
func (e *ConfigError) Unwrap() error {
	return e.Err
}

// Error implements the standard error interface.
func (e *ConfigError) Error() string {
	switch e.Code {
	case ErrUnsupportedType:
		return fmt.Sprintf("unsupported option type: %s", e.TypeString)
	case ErrExecutionFailed:
		if e.Err != nil {
			return fmt.Sprintf("option apply failed [type:%s]: %v", e.TypeString, e.Err)
		}
		return fmt.Sprintf("option apply failed [type:%s]", e.TypeString)
	case ErrEmptyTargetValue:
		return "target for configuration cannot be nil"
	default:
		return "unknown config error"
	}
}

// IsConfigError checks if the given error is a *ConfigError.
func IsConfigError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError)
}

// IsUnsupportedTypeError checks if the error is a ConfigError with the code
// ErrUnsupportedType.
func IsUnsupportedTypeError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrUnsupportedType
}

// IsExecutionFailedError checks if the error is a ConfigError with the code
// ErrExecutionFailed.
func IsExecutionFailedError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrExecutionFailed
}

// IsEmptyTargetValueError checks if the error is a ConfigError with the code
// ErrEmptyTargetValue.
func IsEmptyTargetValueError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Code == ErrEmptyTargetValue
}
