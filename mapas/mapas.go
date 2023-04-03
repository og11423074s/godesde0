package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"

	/*
		fmt.Println(paises)
		fmt.Println(paises["Argentina"])*/

	campeonaro := map[string]int{
		"Barcelona":   39,
		"Real Madrid": 38,
		"Chivas":      37,
		"Boca Junior": 30,
	}

	fmt.Println(campeonaro)

	/*for equipo, puntaje := range campeonaro {
		fmt.Printf("Equipo %s, tiene un puntuaje de %d \n", equipo, puntaje)
	}*/

	delete(campeonaro, "Real Madrid")
	fmt.Println(campeonaro)

	puntaje, existe := campeonaro["Chivas"]
	fmt.Printf("El puntuaje capturado es %d, y el equipo existe = %t\n", puntaje, existe)

}
