package main

import "fmt"

func main() {
	// Ex 1
	Calculator()

	// Ex 2
	count, err := fileLen("/home/brad/copy-to-cubbit.sh")
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Printf("Counted %d bytes\n", count)

	// Ex 3
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))
	fmt.Println(helloPrefix("Maria"))
}
