// Package main demonstrates a Logger configured via configure.Option (Report Q5).
package main

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

func NewLogger(options ...configure.Option[Logger]) (*Logger, error) { // Q5
	l := &Logger{level: "info", out: os.Stdout}
	configure.Apply(l, options)
	if l.level == "" {
		return nil, fmt.Errorf("log level cannot be empty")
	}
	return l, nil
}

func main() {
	logger1, _ := NewLogger()
	fmt.Println("default:", logger1.level)
	logger2, _ := NewLogger(WithLevel("debug"))
	fmt.Println("debug:", logger2.level)
	f, _ := os.CreateTemp("", "log")
	defer os.Remove(f.Name())
	logger3, _ := NewLogger(WithLevel("error"), WithOutput(f))
	fmt.Println("error+file:", logger3.level, logger3.out != nil)
}
