package weirdnums

import (
	"fmt"
	"math"
)

// causes a deliberate overflow for a byte, int and uint
func Overflow() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("Overflow byte: %b, Overflow int32: %d, Overflow uint64: %d\n", b, smallI, bigI)
}
