package main

import "fmt"

func main() {
	// Exemplo de retorno da função Println()
	numBytes, erros1 := fmt.Println("Hello world!", "Oi galera!", 100)
	fmt.Println(numBytes, erros1)

	// Blank identifier representado por _ para não utilizar a variável
	_, erros2 := fmt.Println("Hello world!", "Oi galera!", 100)
	fmt.Println(erros2)

	// Tipos primitivos
	x, y, z := 10, "strings", true
	fmt.Println(x, y, z)
}
