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
			fmt.Println("WaraniApp es una startup en Latinoamerica")
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
