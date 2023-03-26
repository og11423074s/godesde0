package main

import (
	"fmt"

	"github.com/og11423074s/godesde0/variables"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1588)

	fmt.Println(estado)
	fmt.Println(texto)
}
