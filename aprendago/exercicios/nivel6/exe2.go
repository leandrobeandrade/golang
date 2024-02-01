package main

import "fmt"

func Exe2() {
	sli := []int{1, 2, 3, 4}

	fmt.Println()
	fmt.Println(varia_func(sli...))
	fmt.Println(outra_func(sli))
}

func varia_func(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func outra_func(y []int) int {
	total := 0

	for _, v := range y {
		total += v
	}

	return total
}
