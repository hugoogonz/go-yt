package structs

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Credentials struct {
	Username string
	Password string
}

func MyStructs() {
	fmt.Println("Structs in Golang")

	hitesh := User{"Hitesh", "hitesh@gmail.com", true, 25}
	fmt.Println(hitesh)

	users := []User{
		hitesh,
		{Name: "John Doe", Email: "hzdkv@example.com", Status: true, Age: 25},
		{Name: "Tim Burton", Email: "anpch@example.com", Status: true, Age: 25},
		{Name: "Abigail Scotto", Email: "anpch@example.com", Status: true, Age: 25},
	}

	for _, user := range users {
		if user.Name == "Tim Burton" {
			fmt.Printf("%v es un director de cine, productor, escritor y dibujante estadounidense.\n", user.Name)
			break
		}
	}

	credentials := Credentials{
		Username: "wikibaires",
		Password: "contraseña456",
	}
	// Acceder a los campos del struct
	fmt.Println("Username:", credentials.Username) // Imprime: Username: usuario123
	fmt.Println("Password:", credentials.Password) // Imprime: Password: contraseña456

	// Modificar un campo del struct
	credentials.Password = "nuevaContraseña"
	fmt.Println("Nueva Password:", credentials.Password) // Imprime: Nueva Password: nuevaContraseña
}
