package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func Structs() {
	s1 := Employee{
		"Brad",
		"Robinson",
		1,
	}

	s2 := Employee{
		firstName: "Christy",
		lastName:  "Tenney",
		id:        2,
	}

	var s3 Employee
	s3.id = 3
	s3.firstName = "Dr."
	s3.lastName = "Franklin"

	fmt.Println(s1, s2, s3)
}
