package configure

import (
	"fmt"
	"reflect"
)

// Builder provides a fluent interface for collecting and applying options.
// It is ideal for scenarios where configuration options are gathered progressively
// from different parts of an application.
//
// The generic type C represents the configuration type being built, and is
// expected to be a struct type. Using a pointer type for C as the generic
// parameter (e.g., Builder[*MyConfig]) is not recommended as it can lead to
// unexpected behavior and double-pointers.
type Builder[C any] struct {
	opts []any
	base *C
}

// NewBuilder creates a new configuration builder.
// It can optionally take a base configuration object. If provided, this base
// configuration will be cloned and used as the starting point for applying
// options when `Build` is called. If no base is provided, a zero-value
// instance of C will be used.
//
// Panics if the generic type C is itself a pointer type (e.g., Builder[*MyConfig]),
// as this is an unsupported and likely unintended usage pattern that leads to double-pointers.
func NewBuilder[C any](base ...*C) *Builder[C] {
	// Perform the type check at creation time and panic if C is a pointer.
	cType := reflect.TypeOf((*C)(nil)).Elem()
	if cType.Kind() == reflect.Ptr {
		panic(fmt.Sprintf("configure: Builder does not support pointer types for its generic parameter C (e.g., Builder[*MyConfig]), but got %s. Please use a value type like Builder[MyConfig]", cType))
	}

	b := &Builder[C]{}
	if len(base) > 0 && base[0] != nil {
		b.base = base[0]
	}
	return b
}

// Add adds one or more options to the builder. It supports a fluent, chainable API.
func (b *Builder[C]) Add(opts ...any) *Builder[C] {
	b.opts = append(b.opts, opts...)
	return b
}

// AddWhen conditionally adds an option to the builder based on a condition.
// If `condition` is true, `optIfTrue` is added. If `condition` is false and
// `optIfFalse` is provided (as the first element of the variadic parameter),
// then `optIfFalse` is added instead.
// It supports a fluent, chainable API.
func (b *Builder[C]) AddWhen(condition bool, optIfTrue any, optIfFalse ...any) *Builder[C] {
	if condition {
		b.opts = append(b.opts, optIfTrue)
	} else if len(optIfFalse) > 0 {
		b.opts = append(b.opts, optIfFalse[0])
	}
	return b
}

// applyTo applies all collected options to an existing target object.
// This is an internal helper method, as its functionality is exposed via Build() or Apply().
func (b *Builder[C]) applyTo(target *C) (*C, error) {
	return ApplyAny(target, b.opts)
}

// Build creates a new configuration object C and applies all collected options to it.
// It starts with a clone of the base configuration (if set via `NewBuilder`),
// or a zero-value instance of C if no base is provided.
func (b *Builder[C]) Build() (*C, error) {
	// Start with a clone of the base config, or a zero value if no base is set.
	var target C
	if b.base != nil {
		target = *b.base // Create a copy
	}

	// Apply options to the new target.
	return b.applyTo(&target)
}

// Apply implements the ApplierE interface.
// This allows a Builder instance to be passed directly as an option to other
// functions like New or ApplyAny, acting as a "super option".
func (b *Builder[C]) Apply(target *C) error {
	_, err := b.applyTo(target)
	return err
}

// Compile creates a final product `P` by first building a configuration `C`
// using the provided `builder`, and then passing the result to a `factory` function.
// This function acts as the primary top-level entry point for the Config -> Product workflow,
// emphasizing the factory's role in producing the final product from the configuration.
func Compile[C any, P any](factory func(c *C) (*P, error), builder *Builder[C]) (*P, error) {
	config, err := builder.Build()
	if err != nil {
		return nil, newConfigError(ErrExecutionFailed, builder, err)
	}

	return factory(config)
}
