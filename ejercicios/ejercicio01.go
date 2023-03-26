package ejercicios

import (
	"strconv"
	"strings"
)

func ConvierteAEntero(numero string) (int, string) {

	var mensaje string
	numeroInt, _ := strconv.Atoi(strings.TrimSpace(numero))

	if numeroInt > 100 {
		mensaje = "Es mayor a 100"
	} else {
		mensaje = "Es menor a 100"
	}

	return numeroInt, mensaje

}
