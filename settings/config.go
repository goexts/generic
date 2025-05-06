// Package settings implements the functions, types, and interfaces for the module.
package settings

import (
	"os"
)

// FromEnv creates a configuration function that reads from environment variables
func FromEnv[S any, T any](key string, target T) Func[S] {
	return func(s *S) {
		eval := os.Getenv(key)
		if eval != "" {
			//target = T(eval)
		}
	}
}

// Validate creates a configuration function that validates settings
func Validate[S any](validator func(*S)) Func[S] {
	return func(s *S) {
		validator(s)
	}
}

// ValidateE creates a configuration function that validates settings
func ValidateE[S any](validator func(*S) error) FuncE[S] {
	return func(s *S) error {
		if err := validator(s); err != nil {
			return err
		}
		return nil
	}
}
