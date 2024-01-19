package main

import "fmt"

// variável com package level scope
var fora string

func main() {
	// Exemplo de retorno da função Println()
	numBytes, erros1 := fmt.Println("Hello world!", "Oi galera!", 100)
	fmt.Println(numBytes, erros1)

	// Blank identifier representado por _ para não utilizar a variável
	_, erros2 := fmt.Println("Hello world!", "Oi galera!", 100)
	fmt.Println(erros2)

	prin1, prin2 := "Olá", "Mundo!!!"
	sprin := fmt.Sprintln(prin1, prin2) // S... são utilizados para manipular os valores em strings
	fmt.Println(sprin)

	fmt.Println("=== Tipos ===")
	// Tipos primitivos - int, string, boolean
	x, y, z := 10, "strings", true
	fmt.Println(x, y, z)

	fmt.Println("Strings")
	st1 := "Isso é \ninterpretado \tpelo compilador"
	st2 := `"Isso NÃO é \ninterpretado \tpelo compilador"`
	fmt.Println(st1)
	fmt.Println(st2)

	fmt.Println("Conversão de tipos")
	type hotdog int
	var int1 hotdog = 10
	var int2 int = 20
	int2 = int(int1) // Conversão de um tipo criado (hotdog) para outro tipo (int)
	fmt.Println(int2)

	fmt.Println("\n=== Variáveis ===")
	var dentro = "Variável tipada por inferência"
	fora = "Variável decalarada e atribuído um valor"
	fmt.Println(dentro)
	fmt.Println(fora)

	fmt.Println("\n=== Operadores ===")
	// Operador curto de declaração - gopher, funciona passando valores com tipagem por inferência e dentro de code blocks
	a := 10                              // operador de declaração gopher
	a = 20                               // operador de atribuição
	fmt.Printf("a: %v, %d, %T", a, a, a) // imprime valor, valor base 10, tipo
	fmt.Print("\n")
	fmt.Println(a + a)   // operador de soma (40)
	fmt.Println(a - 10)  // operador de subtração (10)
	fmt.Println(a * 10)  // operador de multiplicação (200)
	fmt.Println(a / 10)  // operador de divisão (2)
	fmt.Println(a < 10)  // operador menor (false)
	fmt.Println(a > 10)  // operador maior (true)
	fmt.Println(a == 10) // operador igual (false)
	fmt.Println(a != 10) // operador diferente (true)
	fmt.Println(a >= 10) // operador maior igual (true)
	fmt.Println(a <= 10) // operador menor igual (false)

	//

}
