package loops

import "fmt"

type Product struct {
	Id          string
	ProductName string
	Price       float32
	Country     string
	Status      bool
}

func Loops() {
	fmt.Println("Welcome to loops in Golang")

	countries := []string{"Argentina", "Paraguay", "United Kingdom", "Uruguay", "Mexico", "USA", "Bolivia"}
	fmt.Println(countries)

	for _, c := range countries {
		if c == "Paraguay" {
			fmt.Println("Clasipar es una startup en Latinoamerica")
		}
	}

	for i, v := range countries {
		fmt.Printf("index %v is value %v\n", i, v)
	}

	product := []Product{
		{
			Id:          "1",
			ProductName: "Apple",
			Price:       1000,
			Country:     "Argentina",
			Status:      true,
		},
		{
			Id:          "2",
			ProductName: "Orange",
			Price:       2000,
			Country:     "Uruguay",
			Status:      true,
		},
		{
			Id:          "3",
			ProductName: "Banana",
			Price:       3000,
			Country:     "Paraguay",
			Status:      true,
		},
		{
			Id:          "4",
			ProductName: "Pineapple",
			Price:       4000,
			Country:     "Mexico",
			Status:      true,
		},
	}

	for _, p := range product {
		if p.Country == "Argentina" {
			fmt.Printf("%v avaible in %v \n", p.ProductName, p.Country)
		}
	}
}

// Count of positives / sum of negatives

// Given an array of integers.
// Return an array, where the first element is the count of positives numbers and the second element is sum of negative numbers. 0 is neither positive nor negative.
// If the input is an empty array or is null, return an empty array.

// Example
// For input [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15], you should return [10, -65]

// Link: https://www.codewars.com/kata/576bb71bbbcf0951d5000044

func CountPositivesSumNegatives(numbers []int) []int {
	pos, sum := 0, 0
	for _, num := range numbers {
		if num > 0 {
			pos++
		} else {
			sum += num
		}
	}
	return []int{pos, sum}
}
