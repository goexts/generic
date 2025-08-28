package configure

// Builder provides a fluent interface for collecting and applying options.
// It is ideal for scenarios where configuration options are gathered progressively
// from different parts of an application.
//
// The generic type C represents the configuration type being built.
type Builder[C any] struct {
	opts []any
}

// NewBuilder creates a new, empty configuration builder.
func NewBuilder[C any]() *Builder[C] {
	return &Builder[C]{}
}

// Add adds one or more options to the builder. It supports a fluent, chainable API.
func (b *Builder[C]) Add(opts ...any) *Builder[C] {
	b.opts = append(b.opts, opts...)
	return b
}

// AddWhen conditionally adds an option to the builder if the condition is true.
// It supports a fluent, chainable API.
func (b *Builder[C]) AddWhen(condition bool, opt any) *Builder[C] {
	if condition {
		b.opts = append(b.opts, opt)
	}
	return b
}

// ApplyTo applies all collected options to an existing target object.
func (b *Builder[C]) ApplyTo(target *C) (*C, error) {
	return ApplyAny(target, b.opts)
}

// Build creates a new, zero-value instance of the configuration object C and
// applies all collected options to it.
// The resulting object can then be used directly or passed to a factory.
func (b *Builder[C]) Build() (*C, error) {
	var zero C
	return ApplyAny(&zero, b.opts)
}

// Apply implements the ApplierE interface.
// This allows a Builder instance to be passed directly as an option to other
// functions like New or ApplyAny, acting as a "super option".
func (b *Builder[C]) Apply(target *C) error {
	_, err := b.ApplyTo(target)
	return err
}

// Compile creates a final product `P` by first building a configuration `C`
// using the provided builder, and then passing the result to a factory function.
// This is the primary top-level function for the Config -> Product workflow,
// ensuring a clean separation between configuration and compilation.
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
