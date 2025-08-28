// Package configure provides utilities for applying functional options to objects.
package configure

// Builder provides a fluent interface for collecting and applying options.
// The generic type C represents the configuration type being built.
type Builder[C any] struct {
	opts []any
}

// NewBuilder creates a new configuration builder.
func NewBuilder[C any]() *Builder[C] {
	return &Builder[C]{}
}

// Add adds options to the builder.
func (b *Builder[C]) Add(opts ...any) *Builder[C] {
	b.opts = append(b.opts, opts...)
	return b
}

// AddWhen conditionally adds an option.
func (b *Builder[C]) AddWhen(condition bool, opt any) *Builder[C] {
	if condition {
		b.opts = append(b.opts, opt)
	}
	return b
}

// ApplyTo applies the collected options to the target object.
// It returns an error if any of the options fail.
func (b *Builder[C]) ApplyTo(target *C) (*C, error) {
	return ApplyAny(target, b.opts)
}

// Build applies the collected options to a new zero-value object.
// It returns the configured object, which can then be used by a factory.
func (b *Builder[C]) Build() (*C, error) {
	var zero C
	return ApplyAny(&zero, b.opts)
}

// Apply implements the ApplierE interface, allowing the builder to be used as an option.
func (b *Builder[C]) Apply(target *C) error {
	_, err := b.ApplyTo(target)
	return err
}

// Compile creates a final product `P` by first building a configuration `C`
// using the provided builder, and then passing it to the factory function.
// This is the primary function for the Config -> Product workflow.
func Compile[C any, P any](builder *Builder[C], factory func(c *C) (*P, error)) (*P, error) {
	// First, build the configuration object.
	config, err := builder.Build()
	if err != nil {
		// Wrap the error to provide more context about the configuration build failure.
		return nil, newConfigError(ErrExecutionFailed, builder, err)
	}

	// Then, use the factory to create the final product.
	return factory(config)
}
