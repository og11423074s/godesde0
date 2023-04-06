package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func MiMiddleware() {
	fmt.Println("Inicio")

	result := operacionesMiddle(sumar)(2, 3)
	fmt.Println(result)

	result = operacionesMiddle(restar)(5, 2)
	fmt.Println(result)

	result = operacionesMiddle(multiplicar)(5, 5)
	fmt.Println(result)
}

func operacionesMiddle(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de Operacion....")
		return f(a, b)
	}
}
