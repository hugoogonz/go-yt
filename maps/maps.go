package maps

import "fmt"

func MyMaps() {
	fmt.Println("Maps in Golang")

	languages := make(map[string]string)
	languages["py"] = "Python"
	languages["go"] = "Golang"
	languages["rb"] = "Ruby"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("Go short is: ", languages["go"])

	languagesv2 := map[string]string{
		"py": "Python",
		"go": "Golang",
		"rb": "Ruby",
	}
	fmt.Println("List of all languages: ", languagesv2)

	for k, v := range languages {
		fmt.Printf("For key: %v, value is %v \n", k, v)
	}

	var new_value string
	for key, value := range languagesv2 {
		if key == "go" {
			new_value = value
			break
		}
	}
	fmt.Println("I love it: ", new_value)
}
