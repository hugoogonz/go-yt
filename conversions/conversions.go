package conversions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Conversions() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		panic("Oh no! Something went wrong")
	}

	fmt.Println("Add 1 to your rating: ", numRating+1)
}
