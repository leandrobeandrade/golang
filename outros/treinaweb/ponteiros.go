package main

import (
	"fmt"
)

func mudaValor(a *int) {
	*a = 20
}

func main() {
	var a int = 10

	fmt.Printf("Valor de a: %d\n", a)
	mudaValor(&a)
	fmt.Printf("Valor de a: %d\n", a)
}
