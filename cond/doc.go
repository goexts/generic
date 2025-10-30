/*
Package cond provides generic, ternary-like conditional functions.

This package offers a concise way to express simple conditional logic in a single
expression, avoiding the need for a more verbose if-else block. It provides
functions for both eagerly-evaluated values (If) and lazily-evaluated functions
(IfFunc).

# Usage

## Eager Evaluation with `If`

The `If` function is a direct replacement for a simple ternary operation. Both the
`trueVal` and `falseVal` are evaluated before the function is called.

	// Determine a status message based on an error.
	func GetStatusMessage(err error) string {
		return cond.If(err == nil, "Status: OK", "Status: Failed")
	}

	// Assign a value based on a boolean condition.
	isLoggedIn := true
	buttonText := cond.If(isLoggedIn, "Logout", "Login")

## Lazy Evaluation with `IfFunc`

The `IfFunc` function accepts functions that are only executed when needed. This is
useful when the values to be returned involve expensive computations.

	// The expensive functions are only called when the condition is met.
	func calculateResult(isComplex bool) int {
		return cond.IfFunc(isComplex, doComplexCalculation, doSimpleCalculation)
	}

	func doComplexCalculation() int {
		// Simulate a time-consuming operation.
		time.Sleep(1 * time.Second)
		return 100
	}

	func doSimpleCalculation() int {
		return 1
	}

Choose `If` for simple values and `IfFunc` to optimize performance when dealing
with computationally intensive branches.
*/
package cond
