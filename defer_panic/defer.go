package defer_panic

import (
	"fmt"
	"log"
)

func VamosDefer() {
	fmt.Println("Estre es un primer mensahe")
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic() {

	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero panic \n %v", reco)
		}
	}()

	if a := 1; a == 1 {
		panic("Se encontro el valor 1")
	}
}
