package main

import "fmt"

func Recursividade() {
	r := fatorial_recursao(4)
	l := fatorial_loop(4)

	fmt.Println()
	fmt.Println(r)
	fmt.Println(l)
}

func fatorial_recursao(x int) int {
	if x == 1 {
		return x
	}

	return x * fatorial_recursao(x-1) // x * 3 * 2 * 1 = 24
}

func fatorial_loop(x int) int {
	total := 1

	for x >= 1 {
		total *= x // x * 3 * 2 * 1 = 24
		x--
	}

	return total
}
