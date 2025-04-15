//go:build !race

// Package thread implements the functions, types, and interfaces for the module.
package thread

import (
	"fmt"
	"testing"
)

// Function executes without error and returns expected value in channel
func TestFunctionExecutesWithoutError(t *testing.T) {
	// Arrr! Let's see if this function sails smoothly!
	expected := 42
	f := func() (int, error) {
		return expected, nil
	}
	ch, errCh := AsyncOrErr(f)
	select {
	case v := <-ch:
		if v != expected {
			t.Errorf("Expected %d, but got %d", expected, v)
		}
	case e := <-errCh:
		t.Errorf("Unexpected error: %v", e)
	}
}

// Function handles a successful execution of the provided function
func TestFunctionHandlesSuccessfulExecution(t *testing.T) {
	// Ahoy! Success be on the horizon!
	f := func() (string, error) {
		return "success", nil
	}
	ch, errCh := AsyncOrErr(f)
	select {
	case v := <-ch:
		if v != "success" {
			t.Errorf("Expected 'success', but got %s", v)
		}
	case e := <-errCh:
		t.Errorf("Unexpected error: %v", e)
	}
}

// Channels are properly closed after execution
func TestChannelsAreClosedAfterExecution(t *testing.T) {
	// Avast! Make sure the channels be closed!
	f := func() (int, error) {
		return 0, nil
	}
	ch, errCh := AsyncOrErr(f)
	val := <-ch
	if val != 0 {
		t.Errorf("Expected 0, but got %d", val)
	}
	_, ok1 := <-ch
	_, ok2 := <-errCh
	if ok1 || ok2 {
		t.Error("Channels were not closed properly")
	}
}

// Provided function returns an error
func TestFunctionReturnsError(t *testing.T) {
	// Shiver me timbers! An error be returned!
	f := func() (int, error) {
		return 0, fmt.Errorf("error")
	}
	ch, errCh := AsyncOrErr(f)
	select {
	case <-ch:
		t.Error("Expected no value, but got one")
	case e := <-errCh:
		if e == nil || e.Error() != "error" {
			t.Errorf("Expected 'error', but got %v", e)
		}
	}
}

// Provided function returns a zero value
func TestFunctionReturnsZeroValue(t *testing.T) {
	// Yo-ho-ho! Zero value be returned!
	f := func() (int, error) {
		return 0, nil
	}
	ch, errCh := AsyncOrErr(f)
	select {
	case v := <-ch:
		if v != 0 {
			t.Errorf("Expected 0, but got %d", v)
		}
	case e := <-errCh:
		t.Errorf("Unexpected error: %v", e)
	}
}

// Function handles nil input function gracefully
func TestFunctionHandlesNilInputGracefully(t *testing.T) {
	// Arrr! Nil input be handled with grace!
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Did not expect panic, but got one: %v", r)
		}
	}()
	var f func() (int, error)
	ch, errCh := AsyncOrErr(f)
	select {
	case <-ch:
		t.Error("Expected no value, but got one")
	case e := <-errCh:
		if e == nil {
			t.Error("Expected an error due to nil function")
		}
	}
}
