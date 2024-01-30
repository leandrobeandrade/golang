package main

import "fmt"

func Anonima() {
	x := 10

	func(val int) {
		fmt.Println()
		fmt.Println("O valor de x é:", val)
	}(x) // Executa a função
}
