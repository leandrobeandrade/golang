package main

import "fmt"

func main() {
	basica()
	param(10, 10)
	fmt.Println(retorno())
	fmt.Println(retorno2(true))
	fmt.Println(varia(1, 2, 3, 4))

	Defer()
	Metodo()
	Interfaces()
	Anonima()
	Expressao()
	Retorno()
	Callback()
	Closure()
	Recursividade()
}

func basica() {
	fmt.Println("Função básica")
}

func param(a, b int) {
	fmt.Println("\nFunção com prâmetros")
	fmt.Println(a + b)
}

func retorno() string {
	return "\nFunção com retorno"
}

// Função com mais de um tipo de retorno
func retorno2(x interface{}) (string, bool) {
	if x == "Oi" {
		return "X é uma string", false
	} else if x == true {
		return "X é um booleano", true
	}
	fmt.Println("\n=== Função com mais de um retorno ===")
	return "X não é nenhum destes", false
}

// Função variádica (N° indeterminado de argumentos ou nenhum)
func varia(vals ...int) int {
	total := 0

	for _, v := range vals {
		total += v
	}
	fmt.Println("\nFunção com número indeterminado de argumentos")
	return total
}
