package function

import "fmt"

// Soma1 -> função com parametros que retorna um tipo de valor
func Soma1(a int, b int) int {
	return a + b
}

// Soma2 -> função com parametros que retorna dois tipos de valores
func Soma2(a int, b int) (string, int) {
	return "Soma:", a + b
}

// Soma3 -> função com parametros com retorno nomeado
func Soma3(a int, b int) (resultado int) {
	resultado = a + b
	return
}

// Soma4 -> função variádica, com números indeterminado de parametros
func Soma4(val ...int) int {
	resultado := 0

	for _, v := range val {
		resultado = v
	}
	return resultado
}

// Soma5 -> função que retona outra função
func Soma5() {
	// func() int -> determina o que será retornado, ou seja, retornará uma função com valor do tipo int
	resultado := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v // 2 valores de cinco
		}

		// Retorno de uma função int
		return func() int {
			return res * res // res == 10
		}
	}

	// Execução das duas funções
	fmt.Println(resultado(5, 5)())
}

type Nums struct {
	Num1 int
	Num2 int
}

func (n Nums) Soma6() {
	result := n.Num1 + n.Num2
	fmt.Println("Soma:", result)
}

func (n Nums) Teste() {
	fmt.Println(10)
}
