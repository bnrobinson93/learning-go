package constants

import "fmt"

// declare a const `value` for int/float
// assign to int i & float f
// Print i and f
func Constants() {
	const value = 20

	var i int = value
	var f float64 = value

	fmt.Printf("Integer: %d, Float: %.2f\n", i, f)
}
