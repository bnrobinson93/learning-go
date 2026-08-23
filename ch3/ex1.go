package main

import "fmt"

func Subslices() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sub1 := greetings[:2]  // everything before 2 = 0, 1
	sub2 := greetings[2:4] // from 2 to 4, inclusive = 2, 3, 4
	sub3 := greetings[3:]  // everything after 3 = 4, 5

	fmt.Println(greetings, sub1, sub2, sub3)
}
