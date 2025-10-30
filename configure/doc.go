/*
Package configure provides a robust, type-safe, and flexible implementation
of the Functional Options Pattern for Go. It is designed to simplify the
initialization of complex objects by allowing optional parameters to be passed
in a clean and readable way.

This pattern is ideal for constructors where many parameters are optional, have
sensible defaults, or where you want to avoid a large number of arguments.

# Usage

The core of the pattern involves defining an `Option` type and functions that
return these options.

## Basic Example: Configuring a Logger

Imagine you are creating a `Logger` that can be configured with different logging
levels and output writers. The default is to log at the "info" level to standard
output.

	// 1. Define the object to be configured.
	type Logger struct {
		level string
		out   io.Writer
	}

	// 2. Create functions that return an `Option` for each configurable field.
	func WithLevel(level string) configure.Option[Logger] {
		return configure.OptionFunc[Logger](func(l *Logger) error {
			l.level = level
			return nil
		})
	}

	func WithOutput(w io.Writer) configure.Option[Logger] {
		return configure.OptionFunc[Logger](func(l *Logger) error {
			l.out = w
			return nil
		})
	}

	// 3. Create a constructor that applies the options to a default instance.
	func NewLogger(options ...configure.Option[Logger]) (*Logger, error) {
		// Start with default values.
		l := &Logger{
			level: "info",
			out:   os.Stdout,
		}

		// Apply any provided options.
		configure.ApplyWith(l, options...)

		// You can also perform validation after applying options.
		if l.level == "" {
			return nil, fmt.Errorf("log level cannot be empty")
		}

		return l, nil
	}

Now, you can create loggers with different configurations easily:

	// A default logger (info level, stdout).
	logger1, _ := NewLogger()

	// A debug-level logger.
	logger2, _ := NewLogger(WithLevel("debug"))

	// A logger that writes to a file.
	file, _ := os.Create("app.log")
	logger3, _ := NewLogger(WithLevel("error"), WithOutput(file))

For more advanced usage, including stateful builders and compilation, refer to the
function-specific documentation.
*/
package configure
