package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func Ex1() {
	me := MakePerson("Brad", "Robinson", 32)
	her := MakePersonPointer("Christy", "Tenney", 30)

	fmt.Printf("%v\n%v\n", me, her)
}
