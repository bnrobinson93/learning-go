// Package exerciseone exposes a single function named Add
package exerciseone

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers and returns the result
//
// For more, check out this link about [addition].
//
// [addition]: https://mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
