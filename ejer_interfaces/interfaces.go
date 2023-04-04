package ejer_interfaces

import (
	"fmt"

	"github.com/og11423074s/godesde0/interfaces"
)

func HumanoRespirando(hu interfaces.Humano) {
	hu.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.Sexo())
}

func HumanoVive(hu interfaces.Humano) {
	hu.EstaVivo()
	fmt.Printf("Estoy vivo: %t \n", hu.EstaVivo())
}
