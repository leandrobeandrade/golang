package main

import "fmt"

func If() {
	fmt.Println("\n=== IF padrão ===")
	a := 2
	if a > 1 {
		fmt.Println("Faz algo!")
	}

	fmt.Println("\n=== IF com declaração inline ===")
	if a := 2; a > 1 {
		fmt.Println("Faz algo!")
	}

	if a := 2; a > 1 && 3 > 1 { // Mútiplas condições não precisa de parênteses
		fmt.Println("Faz algo!")
	}

	fmt.Println("\n=== IF com declaração inline ===")
	if 1 > 2 {
		fmt.Println("Algo errado!")
	} else if 1 == 2 {
		fmt.Println("Algo continua errado!")
	} else {
		fmt.Println("Aqui é o certo!")
	}
}
