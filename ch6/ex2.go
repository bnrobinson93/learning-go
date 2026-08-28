package main

import "fmt"

func UpdateSlice(arr []string, input string) {
	arr[len(arr)-1] = input
	fmt.Println(arr)
}

func GrowSlice(arr []string, input string) {
	arr = append(arr, input)
	fmt.Println(arr)
}

func Ex2() {
	arr := []string{"a", "b", "c", "d", "e"} // length = 5
	fmt.Println("before UpdateSlice", arr)
	UpdateSlice(arr, "f") // this will stick because we're updating something at position <5
	fmt.Println("after UpdateSlice", arr)

	fmt.Println("before GrowSlice", arr)
	GrowSlice(arr, "g") // this won't stick because the len has to change
	fmt.Println("after GrowSlice", arr)

	// Cap doesn't matter
	arr2 := make([]string, 1, 2)
	arr2[0] = "a"
	fmt.Println("before GrowSlice with bigger cap", arr2)
	GrowSlice(arr2, "b") // this won't stick because the len has to change
	fmt.Println("after GrowSlice with bigger cap", arr2)
}
