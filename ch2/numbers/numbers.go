package numbers

import "fmt"

// Declare an integer named i with value 20
// Assign i to a float, f
// Print i and f
func IntsAndFloats() {
	var f float64
	var i int = 20

	f = float64(i)

	fmt.Printf("Integer: %d, Float: %.2f\n", i, f)
}
