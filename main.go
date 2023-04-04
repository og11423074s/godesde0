package main

import (
	e "github.com/og11423074s/godesde0/ejer_interfaces"
	"github.com/og11423074s/godesde0/modelos"
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

	/*numeroInt, mensaje := ejercicios.ConvierteAEntero("rrrr ")

	fmt.Println(mensaje)
	fmt.Println(numeroInt)*/

	/*teclado.IngresoNumero()*/

	/*iteraciones.Iterar()*/

	/*fmt.Println(ejercicios.TablaMultiplicar())*/

	//files.GrabaTabla()

	//files.SumaTabla()

	//files.LeoArchivo()

	//funciones.Calculo()

	//funciones.LlamarClosure()

	//funciones.Exponencia(2)

	//arreglos_slice.MuestrSrreglos()

	//arreglos_slice.Capacidad()

	//mapas.MostrarMapas()

	//users.AltaUsuario()

	Pedro := new(modelos.Hombre)
	e.HumanoRespirando(Pedro)

	Maria := new(modelos.Mujer)
	e.HumanoRespirando(Maria)

}
