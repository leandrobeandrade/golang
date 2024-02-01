package main

import "fmt"

func Exe10() {
	clos('A', 5)
	clos('B', 5)
}

func clos(letra rune, tamanho int) {
	lt := letra

	fmt.Println()

	for i := 1; i <= tamanho; i++ {
		fmt.Printf("Valor Armazenado em: [%c]: %d\n", lt, i)
	}
}
