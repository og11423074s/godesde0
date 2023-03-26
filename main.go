package main

import (
	"fmt"

	"github.com/og11423074s/godesde0/ejercicios"
)

func main() {
	//variables.MuestroEnteros()
	//variables.RestoVariables()

	/*estado, texto := variables.ConviertoATexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)*/

	/*
		if os := runtime.GOOS; os == "linux." || os == "OS X." {
			fmt.Println("Esto no es Windows, es ", os)
		} else {
			fmt.Println("Esto es Windows")
		}

		switch os := runtime.GOOS; os {
		case "linux":
			fmt.Println("Esto en Linux")
		case "darwin":
			fmt.Println("Esto en Darwin")
		default:
			fmt.Printf("%s \n", os)
		}*/

	numeroInt, mensaje := ejercicios.ConvierteAEntero("90 ")

	fmt.Println(mensaje)
	fmt.Println(numeroInt)

}
