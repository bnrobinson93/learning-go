package main

import "fmt"

func Ex3() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i // total is a new copy every time. Make sure not to use :=
		fmt.Println(total)
	}
	fmt.Println(total)
}
