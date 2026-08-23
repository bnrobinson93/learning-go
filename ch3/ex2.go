package main

import "fmt"

// Prints the fourth rune as a character, not a number
func Runes() {
	message := "Hi 👩 and 👨"

	runes := []rune(message) // this will have the char code

	fmt.Println(string(runes[3]))
}
