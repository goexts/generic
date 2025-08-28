/*
Package configure provides a robust, type-safe, and flexible implementation
of the Functional Options Pattern for Go. It is designed to handle a wide
range of configuration scenarios, from simple object initialization to complex,
multi-stage product compilation.

# Core Concepts

The package is built around a few core ideas:

  - **Type-Safe Application**: For the highest performance and compile-time safety,
    use the `Apply` and `ApplyE` functions. They are ideal when all options
    are of the same, known type.

    // Example of simple, type-safe configuration:
    type Server struct {
    Port int
    Host string
    }
    type Option func(*Server)
    func WithPort(p int) Option {
    return func(s *Server) { s.Port = p }
    }
    func WithHost(h string) Option {
    return func(s *Server) { s.Host = h }
    }

    server := &Server{}
    configure.Apply(server, []Option{
    WithPort(8080),
    WithHost("localhost"),
    })

  - **Flexible Application**: When you need to handle a mix of different option
    types, use `ApplyAny`. This function uses type assertions to provide
    flexibility, at the cost of compile-time safety and a minor performance overhead.

    opts := []any{
    WithPort(8080),
    func(s *Server) { s.Host = "example.com" }, // A raw function
    }
    server, err := configure.New[Server](opts...)

  - **Stateful Builder**: For scenarios where options are collected progressively
    from different parts of your application, use the `Builder`. It provides a
    fluent, chainable API.

    builder := configure.NewBuilder[Server]().
    Add(WithPort(443)).
    AddWhen(isProduction, WithHost("prod.server.com"))

    server, err := builder.Build()

  - **Compilation**: For the advanced use case of transforming a configuration
    object `C` into a final product `P`, use the top-level `Compile` function.
    This separates the configuration logic from the product creation logic.

    // Example: Using a `ClientConfig` to create an `*http.Client`
    type ClientConfig struct {
    Timeout time.Duration
    }
    factory := func(c *ClientConfig) (*http.Client, error) {
    return &http.Client{Timeout: c.Timeout}, nil
    }

    configBuilder := configure.NewBuilder[ClientConfig]().
    Add(func(c *ClientConfig) { c.Timeout = 20 * time.Second })

    httpClient, err := configure.Compile(configBuilder, factory)

By combining these tools, developers can choose the right approach for their
specific needs, ensuring code remains clean, maintainable, and robust.
*/
package configure
