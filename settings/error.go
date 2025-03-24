// Package settings implements the functions, types, and interfaces for the module.
package settings

import (
	"errors"
	"fmt"
)

// Error types for configuration errors
const (
	ErrUnsupportedType = iota
	ErrExecutionFailed
	ErrEmptyTargetValue
)

type ConfigError struct {
	Type       int    // Error category
	TypeString string // Setting type info
	Err        error  // Original error
}

// 创建错误时记录类型信息
func newConfigError(t int, setting any, err error) *ConfigError {
	return &ConfigError{
		Type:       t,
		TypeString: fmt.Sprintf("%T", setting),
		Err:        err,
	}
}

func (e *ConfigError) Unwrap() error {
	return e.Err
}

// Error message display is enhanced
func (e *ConfigError) Error() string {
	switch e.Type {
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

func IsConfigError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError)
}

func IsUnsupportedTypeError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Type == ErrUnsupportedType
}

func IsEmptyTargetValueError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Type == ErrEmptyTargetValue
}

func IsExecutionFailedError(err error) bool {
	var configError *ConfigError
	return errors.As(err, &configError) && configError.Type == ErrExecutionFailed
}
