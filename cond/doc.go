/*
Package cond provides generic, ternary-like conditional functions.

This package offers a concise way to express simple conditional logic in a single
expression, avoiding the need for a more verbose if-else block. It provides
functions for both eagerly-evaluated values and lazily-evaluated functions.

Example (Eager Evaluation):

	// Instead of:
	// var status string
	// if code == 200 {
	// 	status = "OK"
	// } else {
	// 	status = "Error"
	// }

	status := cond.If(code == 200, "OK", "Error")

Example (Lazy Evaluation):

	// The expensive functions are only called when needed.
	result := cond.IfFunc(isComplex, doComplexCalculation, doSimpleCalculation)
*/
package cond
