package main

import "fmt"

func main() {
	// 1
	doubleMe := 3
	doubled := Doubler(doubleMe)
	fmt.Printf("%d doubled is: %d\n", doubleMe, doubled)

	doubleMeFloat := 2.3
	doubledFloat := Doubler(doubleMeFloat)
	fmt.Printf("%0.2f doubled is: %0.2f\n", doubleMeFloat, doubledFloat)

	// 2
	i := MyIntPrinter(5)
	Print(i)
	f := MyFloatPrinter(2.0)
	Print(f)

	// 3
	myList := List[int]{}
	myList.Add(1)
	myList.Add(2)
	myList.Add(3)
	if err := myList.Insert(4, 10); err != nil {
		fmt.Println(err)
	}
	if err := myList.Insert(4, 1); err != nil {
		fmt.Println(err)
		fmt.Println(err)
	}
	if err := myList.Insert(5, 0); err != nil {
		fmt.Println(err)
	}
	if err := myList.Insert(6, 3); err != nil {
		fmt.Println(err)
	}
	myList.Add(7)
	fmt.Println("Location of '2': ", myList.Index(2))
	fmt.Println("Location of '9': ", myList.Index(9))
	myList.Print()
}
