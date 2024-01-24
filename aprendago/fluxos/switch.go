package main

import "fmt"

func Switch() {
	fmt.Println("\n=== Switch padrão ===")
	x := 5
	switch {
	case x == 5:
		fmt.Println("X igual a 5") // Para aqui na primeira ocorrência verdadeira
	case x < 6:
		fmt.Println("X menor que 5") // Mesmo verdadeiro parou na primeira encontrada
	case x > 5:
		fmt.Println("X maior que 5")
	}

	fmt.Println("\n=== Switch com declaração default ===")
	y := "Fulano"
	switch y {
	case "Ciclano":
		fmt.Println("Nome é Ciclano")
	case "Beltrano":
		fmt.Println("Nome é Beltrano")
	case "Fulano":
		fmt.Println("Nome é Fulano")
	case "Belciclano":
		fmt.Println("Nome é Belciclano")
	}
}
