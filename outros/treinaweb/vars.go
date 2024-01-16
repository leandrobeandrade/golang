package main

import "fmt"

var global string = "Olá mundo GO >>>"

func vars() {
	fmt.Println("===== VARIAVEIS =====")

	// variável disponível em todo escopo
	fmt.Println(global)

	// declaração com inferência do tipo
	var idade1 = 10
	fmt.Println("tipo int inferido ", idade1)

	// declaração e atribuição direta
	var idade2 int = 18
	fmt.Println("tipada com valor ", idade2)

	// declaração tipada com atribuição posterior
	var idade3 int
	idade2 = 30

	fmt.Println("atribuição posterior ", idade3)

	// operador gopher só funciona em funções e com inferência do tipo
	idade4 := 15
	fmt.Println("gopher", idade4)

	// constantes devem ser inicializadas e não precisam ser tipadas
	const PI float64 = 3.141592653589793
	fmt.Println("const", PI)
}
