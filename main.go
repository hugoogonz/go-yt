package main

import (
	"fmt"

	"github.com/ubuntu-bros/go-yt/arrays"
	"github.com/ubuntu-bros/go-yt/conversions"
	"github.com/ubuntu-bros/go-yt/functions"
	"github.com/ubuntu-bros/go-yt/hello"
	"github.com/ubuntu-bros/go-yt/ifelse"
	"github.com/ubuntu-bros/go-yt/loops"
	"github.com/ubuntu-bros/go-yt/maps"
	"github.com/ubuntu-bros/go-yt/mytime"
	"github.com/ubuntu-bros/go-yt/pointers"
	"github.com/ubuntu-bros/go-yt/slices"
	"github.com/ubuntu-bros/go-yt/structs"
	"github.com/ubuntu-bros/go-yt/switchcase"
	"github.com/ubuntu-bros/go-yt/userinput"
	"github.com/ubuntu-bros/go-yt/variables"
)

func main() {
	fmt.Println("Hey! Ubuntu Bros!")

	// 01-hello
	hello.Hello()
	// 02-variable
	variables.Variables()
	// 03-userinput
	userinput.UserInput()
	// 04-conversions
	conversions.Conversions()
	// 05-mytime
	mytime.MyTime()
	// 06-pointers
	pointers.Pointers()
	// 07-arrays
	arrays.MyArray()
	// 08-slices
	slices.MySlices()
	// 09-maps
	maps.MyMaps()
	// 10-structs
	structs.MyStructs()
	// 11-ifelse
	ifelse.IfElse()
	// 12-switchcase
	switchcase.SwitchCase()
	// 13-loops
	loops.Loops()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	result := loops.CountPositivesSumNegatives(arr)
	fmt.Println(result)

	// 14-functions
	fmt.Println(functions.MakeNegative(5))
	fmt.Println(functions.MakeNegative(-5))

	// 15-methods

}
