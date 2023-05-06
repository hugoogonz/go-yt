package arrays

import "fmt"

func MyArray() {
	fmt.Println("Welcome to array in Golang")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[2] = "Banana"
	fruitList[3] = "Mango"
	fmt.Println("Fruit list is: ", fruitList)      // [Apple Orange Banana Mango]
	fmt.Println("Fruit list is: ", len(fruitList)) // 4

	vegList := []string{"Tomato", "Lettuce", "Onion", "Carrot"}
	fmt.Println("Veg list is: ", vegList)      // [Tomato Lettuce Onion Carrot]
	fmt.Println("Veg list is: ", len(vegList)) // 4

	for i := 0; i < len(vegList); i++ {
		if vegList[i] == "Carrot" {
			fmt.Println("I love it!")
		}
	}
}
