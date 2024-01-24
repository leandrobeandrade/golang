package main

import "fmt"

func Exe6() {
	x := 5
	switch {
	case x == 5:
		fmt.Println("X igual a 5") // Para aqui na primeira ocorrência verdadeira (no fall-through)
	case x < 6:
		fmt.Println("X menor que 6") // Mesmo verdadeiro parou na primeira ocorrência encontrada
	case x > 5:
		fmt.Println("X maior que 5")
	}
}
