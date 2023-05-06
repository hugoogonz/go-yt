package variables

import "fmt"

var Name string
var Country string

var jwtToken string

func Variables() {
	Name = "John Doe"
	Country = "Germany"
	fmt.Println(Name, Country)

	jwtToken = "949d1u22cbffbrarjh182eig55721odj"
	fmt.Println(jwtToken)
	fmt.Printf("Variable is of type: %T\n", jwtToken)

	name := "Tortugita"
	country := "Argentina"
	fmt.Println(name, country)

	var username string = "wikipedia"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T\n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T\n", smallVal)

	var smallFloat float32 = 1537.55
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T\n", smallFloat)

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T\n", anotherVariable)

	// implicit type
	var website = "http://www.github.com"
	fmt.Println(website)
}
