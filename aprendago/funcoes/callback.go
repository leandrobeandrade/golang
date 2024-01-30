package main

import "fmt"

// Função variádica que retorna inteiro
func soma(x ...int) int {
	n := 0

	for _, v := range x {
		n += v
	}

	return n
}

// SomentePares - recebe uma outra função e um slice de inteiros como argumento e retorna inteiro
// Soma - recebe valores pares e é executada
func SomentePares(f func(...int) int, y ...int) int {
	var sli []int

	for _, v := range y {
		if v%2 == 0 {
			sli = append(sli, v)
		}
	}

	total := f(soma(sli...))
	return total
}

func Callback() {
	pares := SomentePares(soma, []int{1, 2, 3, 4}...)
	fmt.Println()
	fmt.Println("Soma dos números pares:", pares)
	fmt.Println()
}
