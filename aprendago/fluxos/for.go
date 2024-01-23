package main

import "fmt"

func For() {
	fmt.Println("=== FOR comum ===")

	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}

	fmt.Println("\n=== FOR separado ===")

	i := 1
	for i < 6 {
		fmt.Println(i)
		i++
	}

	/* fmt.Println("=== FOR infinito ===")

	for {
		fmt.Println("Infinito")
	} */

	fmt.Println("\n=== FOR com BREAK ===")

	j := 1
	for {
		if j < 6 {
			fmt.Println("Incrementando!! J ==", j)
			j++
		} else {
			fmt.Println("Encerrou o loop no:", i)
			break // encerra o loop
		}
	}

	fmt.Println("\n=== FOR com CONTINUE ===")

	for i := 0; i <= 20; i++ {
		if i%2 != 0 {
			continue // pula e continua para o próxima iteração
		}
		fmt.Println("Pares", i)
	}
}
