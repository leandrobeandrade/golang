package main

import (
	"fmt"
)

func main() {
	x := 1
	y := &x // ponteiro de x

	fmt.Print("Valor de x: ", x, " Valor de y: ", y)

	*y++ // altera valor de x pelo ponteiro y

	fmt.Println("\nValor de x:", *y)
	fmt.Printf("\n%T, %T\n", x, y)
	fmt.Println(x, y)
	fmt.Println()

	a := 9
	b := 9

	fmt.Println("Valor original de a:", a)
	fmt.Println("Valor original de b:", b)

	estarecebeovalor(a)
	estarecebeumponteiro(&b)

	fmt.Println("Valor original de a:", a)
	fmt.Println("Valor original de b:", b)
}

// Muda o valor do parametro da função, mas o valor original continua o mesmo
func estarecebeovalor(vlr int) {
	vlr++
	fmt.Println("Valor de a alterado pela função:", vlr)
}

// Muda o valor original pelo ponteiro
func estarecebeumponteiro(vlr *int) {
	*vlr++
	fmt.Println("Valor de b alterado pela função:", *vlr)
}
