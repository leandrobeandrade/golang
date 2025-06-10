package funcs

import "fmt"

func Somar(a int, b int) int {
	return a + b
}

func Subtrair(a int, b int) int {
	return a - b
}

func Multiplicar(a int, b int) int {
	return a * b
}

func Dividir(a int, b int) int {
	return a / b
}

func AllFuncs() {
	fmt.Println(Somar(10, 20))
	fmt.Println(Subtrair(10, 20))
	fmt.Println(Multiplicar(10, 20))
	fmt.Println(Dividir(10, 20))
}
