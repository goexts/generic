// Package examples designs a Connection constructor with dependent options (Report Q10).
// Extended example: covers three option forms (alias, defined type, plain function), multiple Apply usages, interface-style options, and error-returning options.
package examples

import (
	"errors"
	"fmt"
	"time"

	"github.com/goexts/generic/configure"
)

type Connection struct {
	retries    int
	retryDelay time.Duration
}

// OptionA alias type for configure.Option[Connection]
type OptionA = configure.Option[Connection]

// OptionB custom named type (same underlying type as OptionA), used to demonstrate custom option style
type OptionB configure.Option[Connection]

// OptionC plain function type, independent of external packages
type OptionC func(*Connection)

// WithRetries A-style options (alias)
func WithRetries(n int) OptionA              { return func(c *Connection) { c.retries = n } }
func WithRetryDelay(d time.Duration) OptionA { return func(c *Connection) { c.retryDelay = d } }

// WithRetriesFunc B-style options (named type)
func WithRetriesFunc(n int) OptionB              { return func(c *Connection) { c.retries = n } }
func WithRetryDelayFunc(d time.Duration) OptionB { return func(c *Connection) { c.retryDelay = d } }

func ExampleConnection() {
	conn := &Connection{}
	WithRetries(3)(conn)
	WithRetryDelay(100 * time.Millisecond)(conn)
	fmt.Println("Connection:", conn.retries, conn.retryDelay)

	// Output:
	// Connection: 3 100ms
}

// WithRetriesC C-style options (plain function type)
func WithRetriesC(n int) OptionC              { return func(c *Connection) { c.retries = n } }
func WithRetryDelayC(d time.Duration) OptionC { return func(c *Connection) { c.retryDelay = d } }

// ToOptionA converts OptionC to OptionA for reuse
func ToOptionA(o OptionC) OptionA { return func(c *Connection) { o(c) } }
func ToOptionB(o OptionC) OptionB { return func(c *Connection) { o(c) } }

// Demonstration of multiple Apply usages
// 1) Use configure.Apply to apply OptionA
// 2) Use configure.New to construct and apply OptionB
// 3) Manually loop and apply OptionC

// NewConnectionA constructs a Connection using OptionA and applies to configure.Apply.
// NewConnectionA enforces that retryDelay is only valid if retries > 0.
func NewConnectionA(options ...OptionA) (*Connection, error) { // Q10
	c := &Connection{retries: 0, retryDelay: 0}
	configure.Apply(c, options)
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

// NewConnectionB constructs a Connection using OptionB and applies to configure.New.
// NewConnectionB enforces the same constraint as NewConnectionA.
func NewConnectionB(options ...OptionB) (*Connection, error) { // Q10
	c := configure.New[Connection](options)
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

// NewConnectionC constructs a Connection using OptionC applied by manual loop.
func NewConnectionC(options ...OptionC) (*Connection, error) {
	c := &Connection{}
	for _, opt := range options {
		opt(c)
	}
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

// ConnOption interface-style options: suitable for more complex option objects
type ConnOption interface{ Apply(*Connection) }

// SetRetries sets the retries count for Connection.
type SetRetries struct{ N int }

func (o SetRetries) Apply(c *Connection) { c.retries = o.N }

type SetRetryDelay struct{ D time.Duration }

func (o SetRetryDelay) Apply(c *Connection) { c.retryDelay = o.D }

// NewConnectionI constructs a Connection using interface-style options.
func NewConnectionI(options ...ConnOption) (*Connection, error) {
	c := &Connection{}
	for _, opt := range options {
		opt.Apply(c)
	}
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

// EOption is a functional option that returns an error
type EOption func(*Connection) error

func WithRetriesErr(n int) EOption {
	return func(c *Connection) error {
		if n < 0 {
			return errors.New("retries must be >= 0")
		}
		c.retries = n
		return nil
	}
}

func WithRetryDelayErr(d time.Duration) EOption {
	return func(c *Connection) error {
		if d < 0 {
			return errors.New("retryDelay must be >= 0")
		}
		c.retryDelay = d
		return nil
	}
}

// ConnOptionErr applies configuration to Connection and may return an error
type ConnOptionErr interface{ Apply(*Connection) error }

// SetRetriesErr sets the retries count for Connection.
type SetRetriesErr struct{ N int }

// Apply implements ConnOptionErr
func (o SetRetriesErr) Apply(c *Connection) error {
	if o.N < 0 {
		return errors.New("retries must be >= 0")
	}
	c.retries = o.N
	return nil
}

// SetRetryDelayErr sets the retryDelay for Connection.
type SetRetryDelayErr struct{ D time.Duration }

// Apply implements ConnOptionErr
func (o SetRetryDelayErr) Apply(c *Connection) error {
	if o.D < 0 {
		return errors.New("retryDelay must be >= 0")
	}
	c.retryDelay = o.D
	return nil
}

// NewConnectionErr constructs a Connection using error-returning functional options
func NewConnectionErr(options ...EOption) (*Connection, error) {
	c := &Connection{}
	for _, opt := range options {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

// NewConnectionIErr constructs a Connection using error-returning interface options
func NewConnectionIErr(options ...ConnOptionErr) (*Connection, error) {
	c := &Connection{}
	for _, opt := range options {
		if err := opt.Apply(c); err != nil {
			return nil, err
		}
	}
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

func ExampleNewConnectionA() {
	// 1) OptionA + configure.Apply
	okA, _ := NewConnectionA(WithRetries(3), WithRetryDelay(200*time.Millisecond))
	fmt.Println("okA:", okA.retries, okA.retryDelay)

	// 2) OptionB + configure.New (recommended: function type OptionB + NewConnectionB)
	okB, _ := NewConnectionB(WithRetriesFunc(2), WithRetryDelayFunc(100*time.Millisecond))
	fmt.Println("okB:", okB.retries, okB.retryDelay)

	// 3) OptionC + manual application
	okC, _ := NewConnectionC(WithRetriesC(4), WithRetryDelayC(300*time.Millisecond))
	fmt.Println("okC:", okC.retries, okC.retryDelay)

	// 4) Adaptation: adapt C-style options to A/B and use their constructors
	okA2, _ := NewConnectionA(ToOptionA(WithRetriesC(5)), ToOptionA(WithRetryDelayC(150*time.Millisecond)))
	fmt.Println("okA2:", okA2.retries, okA2.retryDelay)
	okB2, _ := NewConnectionB(ToOptionB(WithRetriesC(1)), ToOptionB(WithRetryDelayC(50*time.Millisecond)))
	fmt.Println("okB2:", okB2.retries, okB2.retryDelay)

	// 5) Interface-style options
	okI, _ := NewConnectionI(SetRetries{N: 2}, SetRetryDelay{D: 75 * time.Millisecond})
	fmt.Println("okI:", okI.retries, okI.retryDelay)

	// 6) Error-returning functional options
	okE, errE := NewConnectionErr(WithRetriesErr(3), WithRetryDelayErr(250*time.Millisecond))
	fmt.Println("okE:", okE.retries, okE.retryDelay, "errE:", errE)

	// 7) Error-returning interface options
	okIE, errIE := NewConnectionIErr(SetRetriesErr{N: 0}, SetRetryDelayErr{D: 400 * time.Millisecond})
	fmt.Println("okIE:", okIE.retries, okIE.retryDelay, "errIE:", errIE)

	// 8) Example violating the constraint: setting retryDelay only will return an error
	_, errBad := NewConnectionA(WithRetryDelay(1 * time.Second))
	fmt.Println("bad err:", errBad != nil)
}
