package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

func Print[T Printable](str T) {
	fmt.Println(str)
}

type MyIntPrinter int

func (i MyIntPrinter) String() string {
	return strconv.Itoa(int(i))
}

type MyFloatPrinter float64

func (f MyFloatPrinter) String() string {
	return fmt.Sprintf("%f\n", f)
}
