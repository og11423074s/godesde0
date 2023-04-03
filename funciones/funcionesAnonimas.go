package funciones

func Calculo() {

	var numero3 int = 32
	var numero4 int = 243

	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	println(calculo(10, 25))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}

	println(calculo(10, 25))
}
