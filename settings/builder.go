// Package settings implements the functions, types, and interfaces for the module.
package settings

// Builder defines the interface for building configuration settings
type Builder[S any] interface {
	Add(setting Func[S]) Builder[S]
	AddWhen(condition bool, setting Func[S]) Builder[S]
	Build() []Func[S]
}

// builder implements the Builder interface
type builder[S any] struct {
	settings []Func[S]
}

// NewBuilder creates a new configuration builder
func NewBuilder[S any]() Builder[S] {
	return &builder[S]{}
}

// Add appends a configuration function to the builder
func (b *builder[S]) Add(setting Func[S]) Builder[S] {
	b.settings = append(b.settings, setting)
	return b
}

// AddWhen conditionally adds a configuration function
func (b *builder[S]) AddWhen(condition bool, setting Func[S]) Builder[S] {
	if condition {
		return b.Add(setting)
	}
	return b
}

// Build returns the collected configuration functions
func (b *builder[S]) Build() []Func[S] {
	return b.settings
}

// ApplyBuilder applies settings from a builder
func ApplyBuilder[S any](target *S, builder Builder[S]) *S {
	return Apply(target, builder.Build())
}
