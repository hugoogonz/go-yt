package userinput

import (
	"bufio"
	"fmt"
	"os"
)

func UserInput() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our Pizza: ")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of this rating is %T \n", input)
}
