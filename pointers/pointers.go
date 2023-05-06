package pointers

import "fmt"

func Pointers() {
	// https://www.geeksforgeeks.org/pointers-in-golang/
	fmt.Println("Welcome to a class on pointers")

	// var mypointer *int
	// fmt.Println("Value of mypointer is:", mypointer)

	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of ptr is:", ptr)
	fmt.Println("Value of *ptr is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is:", myNumber)
}
