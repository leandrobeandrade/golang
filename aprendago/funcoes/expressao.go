package main

import "fmt"

func Expressao() {
	x := 20

	y := func(val int) {
		fmt.Println("O valor de x é:", val)
	}

	y(x)
}
