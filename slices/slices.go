package slices

import "fmt"

func MySlices() {
	fmt.Println("Welcome to slices in Golang")

	vegList := []string{"Tomato", "Lettuce", "Onion", "Carrot"}
	fmt.Printf("Type of Veg list is: %T \n", vegList)

	vegList = append(vegList, "Corn")

	// how to remove a value from slices based on index
	fmt.Println(vegList[1:3]) // [Lettuce Onion]
	fmt.Println(vegList[0:1]) // [Tomato]
	fmt.Println(vegList[0:2]) // [Tomato Lettuce]

	coursesList := []string{"Python", "Golang", "Vue", "Gin Gonic", "SolidJs"}
	fmt.Println(coursesList)
	fmt.Println(coursesList[:2]) // [Python Golang]
	fmt.Println(coursesList[2:]) // [Python Golang]

	// Obtener el índice correspondiente a "Vue"
	startIndex := -1
	for i, course := range coursesList {
		if course == "Vue" {
			startIndex = i
			break
		}
	}
	// Imprimir los elementos desde "Vue" hasta el final
	if startIndex != -1 {
		fmt.Println(coursesList[startIndex:])
	}
}
