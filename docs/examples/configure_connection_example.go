// Package main designs a Connection constructor with dependent options (Report Q10).
package main

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

type Option = configure.Option[Connection]

type FuncOption func(*Connection)

func WithRetries(n int) Option { return func(c *Connection) { c.retries = n } }
func WithRetryDelay(d time.Duration) Option {
	return func(c *Connection) { c.retryDelay = d }
}

// NewConnection enforces: retryDelay only valid if retries > 0.
func NewConnection(options ...Option) (*Connection, error) { // Q10
	c := &Connection{retries: 0, retryDelay: 0}
	configure.Apply(c, options)
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

// NewConnectionFunc enforces: retryDelay only valid if retries > 0.
func NewConnectionFunc(options ...FuncOption) (*Connection, error) { // Q10
	c := &Connection{retries: 0, retryDelay: 0}
	for _, opt := range options {
		opt(c)
	}
	if c.retryDelay > 0 && c.retries <= 0 {
		return nil, errors.New("WithRetryDelay requires WithRetries > 0")
	}
	return c, nil
}

func WithRetriesFunc(n int) FuncOption { return func(c *Connection) { c.retries = n } }
func WithRetryDelayFunc(d time.Duration) FuncOption {
	return func(c *Connection) { c.retryDelay = d }
}

func main() {
	ok, _ := NewConnection(WithRetries(3), WithRetryDelay(200*time.Millisecond))
	fmt.Println("ok:", ok.retries, ok.retryDelay)
	// 推荐：使用函数类型选项 FuncOption 搭配 NewConnectionFunc
	ok2, _ := NewConnectionFunc(WithRetriesFunc(2), WithRetryDelayFunc(100*time.Millisecond))
	fmt.Println("ok2:", ok2.retries, ok2.retryDelay)
	_, err := NewConnection(WithRetryDelay(1 * time.Second))
	fmt.Println("bad err:", err != nil)
}
