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

// Função que recebe uma outra função como argumento e retorna inteiro
// Recebe valores, armazena os pares e soma estes valores
func SomentePares(f func (...int) int, y ...int) int {
	var sli []int

	for _, v := range y {
		if v % 2 ==0 {
			sli = append(sli, v)
		}
	}

	total := f(soma(sli...)) 
	return total
}

func Callback()  {
	pares := SomentePares(soma, []int{1,2,3,4}...)
	fmt.Println(pares)
}