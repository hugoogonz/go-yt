package conversions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Conversions() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(input, 64)

	if err != nil {
		panic("Oh no! Something went wrong")
	}

	fmt.Println("Thanks for rating", input)

	fmt.Println(numRating, err)
}
