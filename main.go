package main

import (
	"fmt"

	"github.com/ubuntu-bros/go-yt/arrays"
	"github.com/ubuntu-bros/go-yt/conversions"
	"github.com/ubuntu-bros/go-yt/hello"
	"github.com/ubuntu-bros/go-yt/ifelse"
	"github.com/ubuntu-bros/go-yt/maps"
	"github.com/ubuntu-bros/go-yt/mytime"
	"github.com/ubuntu-bros/go-yt/pointers"
	"github.com/ubuntu-bros/go-yt/slices"
	"github.com/ubuntu-bros/go-yt/structs"
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
}
