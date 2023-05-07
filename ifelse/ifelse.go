package ifelse

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func IfElse() {
	users := []User{
		{Name: "Robert Pattinson", Email: "anpch@example.com", Status: true, Age: 25},
		// {Name: "John Doe", Email: "hzdkv@example.com", Status: true, Age: 25},
	}

	for _, user := range users {
		if user.Name == "Robert Pattinson" {
			fmt.Printf("Robert Douglas Thomas Pattinson es un actor, modelo, productor y cantante británico.\n")
		} else {
			fmt.Println("N/A", user.Name)
		}
	}
}
