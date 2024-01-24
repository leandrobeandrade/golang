package main

import "fmt"

func Switch() {
	fmt.Println("\n=== Switch padrão ===")
	x := 5
	switch {
	case x == 5:
		fmt.Println("X igual a 5") // Para aqui na primeira ocorrência verdadeira (no fall-through)
	case x < 6:
		fmt.Println("X menor que 5") // Mesmo verdadeiro parou na primeira ocorrência encontrada
	case x > 5:
		fmt.Println("X maior que 5")
	}

	fmt.Println("\n=== Switch com valor default ===")
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

	fmt.Println("\n=== Switch com fall-through ===")
	switch y {
	case "Fulano":
		fmt.Println("Nome é Fulano")
		fallthrough // Passa para a próxima ocorrência, mesmo sendo falsa
	case "Ciclano":
		fmt.Println("Nome é Ciclano") // Chega até aqui, pula o case e vai para o statement
	case "Beltrano":
		fmt.Println("Nome é Beltrano")
	case "Belciclano":
		fmt.Println("Nome é Belciclano")
	}

	fmt.Println("\n=== Switch com default case ===")
	switch y {
	case "Ciclano":
		fmt.Println("Nome é Ciclano")
	case "Beltrano":
		fmt.Println("Nome é Beltrano")
	case "Belciclano":
		fmt.Println("Nome é Belciclano")
	default:
		fmt.Println("Não encontrou nenhum!!!")
	}

	fmt.Println("\n=== Switch com case composto ===")
	switch y {
	case "Beltrano", "Zezinho":
		fmt.Println("Time 1")
	case "Fulano", "Ciclano": // Retorna se encontrar uma ocorrência verdadeira
		fmt.Println("Time 2")
	case "Belciclano":
		fmt.Println("Nome é Belciclano")
	default:
		fmt.Println("Não encontrou nenhum!!!")
	}
}
