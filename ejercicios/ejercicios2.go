package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {

	scanner := bufio.NewScanner(os.Stdin)
	var numero int
	var err error
	var texto string

	for {
		fmt.Print("Ingrese un numero: ")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	for i := 1; i <= 10; i++ {
		linea := strconv.Itoa(numero) + " x " + strconv.Itoa(i)
		texto += fmt.Sprintf("%s = %d\n", linea, numero*i)
	}

	return texto

}
