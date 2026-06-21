// Package math contains functions for performing to advanced mathematical operations.
package math

import (
	"golang.org/x/exp/constraints"
)

// Number is a constraint allowing any number type, both integers and floats.
type Number interface {
	constraints.Integer | constraints.Float
}

// Add performs [addition] on two numbers and returns the result.
//
// It works on any type of numbers. Both numbers must be of the same type.
// [addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](x, y T) T {
	return x + y
}
