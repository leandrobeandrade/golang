package main

import (
	"fmt"
)

func main() {
	x := 0
	y := &x // ponteiro de x

	fmt.Print("Valor de x: ", x, " Valor de y: ", y)

	*y++ // altera valor de x pelo ponteiro y

	fmt.Println("\nValor de x:", *y)
	fmt.Printf("\n%T, %T\n", x, y)
	fmt.Print(x, y)
}
