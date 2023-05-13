package methods

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Este método toma una instancia de User (u)
func (u User) GetStatus() {
	// fmt.Println("Is User active: ", u.Status)
	if u.Status == true {
		fmt.Println("User is active")
	} else {
		fmt.Println("User is inactive")
	}
}

func (u User) GetUser() {
	if u.Status == true {
		fmt.Printf("%v is active: \n", u.Name)
	} else {
		fmt.Printf("%v is inactive: \n", u.Name)
	}
}

func main() {

	hitesh := User{Name: "Hitesh", Email: "hzdkv@example.com", Status: true, Age: 20}
	madhur := User{Name: "Madhur", Email: "hzdkv@example.com", Status: false, Age: 20}
	hitesh.GetStatus()
	madhur.GetStatus()

	hitesh.GetUser()
	madhur.GetUser()
}
