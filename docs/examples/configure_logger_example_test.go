// Package examples demonstrates a Logger configured via configure.Option (Report Q5).
package examples

import (
	"fmt"
	"os"

	"github.com/goexts/generic/configure"
)

type Logger struct {
	level string
	out   *os.File
}

func WithLevel(level string) configure.Option[Logger] {
	return func(l *Logger) { l.level = level }
}

func WithOutput(w *os.File) configure.Option[Logger] {
	return func(l *Logger) { l.out = w }
}

func NewLogger(opts ...configure.Option[Logger]) (*Logger, error) {
	logger := &Logger{level: "info", out: os.Stdout}
	for _, opt := range opts {
		opt(logger)
	}
	return logger, nil
}

func ExampleNewLogger() {
	logger1, _ := NewLogger()
	fmt.Println("Default logger level:", logger1.level)

	logger2, _ := NewLogger(WithLevel("debug"))
	fmt.Println("Custom logger level:", logger2.level)

	f, _ := os.CreateTemp("", "log")
	defer os.Remove(f.Name())
	logger3, _ := NewLogger(WithLevel("error"), WithOutput(f))
	fmt.Println("error+file:", logger3.level, logger3.out != nil)

	// Output:
	// Default logger level: info
	// Custom logger level: debug
	// error+file: error true
}
