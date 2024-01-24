package main

import "fmt"

func Exe1() {
	// Números de 1 - 10000
	/* for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	} */

	for i := 65; i < 90; i++ {
		fmt.Println(i)
		for j := 0; j <= 2; j++ {
			fmt.Printf("\t%U '%c'\n ", i, i)
		}
	}
}
