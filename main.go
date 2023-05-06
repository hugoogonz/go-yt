package main

import (
	"fmt"

	"github.com/ubuntu-bros/go-yt/hello"
	"github.com/ubuntu-bros/go-yt/variables"
)

func main() {
	fmt.Println("Hey! Ubuntu Bros!")

	// 01-hello
	hello.Hello()
	// 02-variable
	variables.Variables()
}
