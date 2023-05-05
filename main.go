package main

import (
	"fmt"

	"github.com/ubuntu-bros/go-yt/variables"
)

func main() {
	fmt.Println("Hey! Ubuntu Bros!")
	variables.MuestroEnteros()
	variables.RestoVariables()

	estado, texto := variables.ConviertoATexto(1538)
	fmt.Println(estado, texto)
}
