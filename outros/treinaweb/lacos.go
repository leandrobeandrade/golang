package main

import "fmt"

func lacos() {
	fmt.Println("\n===== LAÇOS =====")

	idade1, idade2, idade3 := 20, 35, 10
	bol := true

	fmt.Println("\n### if/else ###")
	if idade1 >= 18 && true {
		fmt.Println("Maior! ", idade2)

		if bol {
			fmt.Println("Verdadeiro!")
		}
	} else {
		fmt.Println("Menor! ", idade3)
	}

	fmt.Println("\n### for ###")
	val1 := 5
	val2 := []int{0, 1, 2, 3, 4, 5}

	// For comum
	for i := 0; i <= val1; i++ {
		fmt.Println("FOR 1: i =>", i)
	}

	fmt.Println()

	// For com range de valor onde _ é o índice e vl é o valor
	for _, vl := range val2 {
		fmt.Println("FOR 2: i =>", vl)
	}
}
