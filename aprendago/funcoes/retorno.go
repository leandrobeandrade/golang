package main

import "fmt"

func Retorno() {
	x := retornando_funcao(10)
	y := x(2)

	// OU
	z := retornando_funcao(10)(2)

	fmt.Println("O valor do retorno é:", y)
	fmt.Println("O valor do retorno é:", z)
}

// Retorna uma função
func retornando_funcao(x int) func(int) int {
	return func(i int) int {
		return 10 + x*i // 10 + (10 X 2) = 30
	}
}
