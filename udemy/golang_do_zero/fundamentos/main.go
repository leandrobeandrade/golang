package main

import (
	"fmt"
	"udemy/golang_do_zero/fundamentos/pacotes"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("=== Pacotes")
	fmt.Println("Escrevendo do arquivo main")
	pacotes.Escrever()

	res := checkmail.ValidateFormat("tes")
	fmt.Println("ERRO:", res)

	fmt.Println("\n=== Variáveis")
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

	fmt.Println("\n=== Tipos de dados")
	var numero int16 = -10000
	fmt.Println(numero)

	var numero2 int16 = 127
	fmt.Println(numero2)

	var numero3 int16 = -128
	fmt.Println(numero3)

	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3
	fmt.Println(meuArray)

	meuArray2 := [3]int{1, 2, 3}
	fmt.Println(meuArray2)

	meuSlice := []int{1, 2, 3}
	fmt.Println(meuSlice)

	meuSlice = append(meuSlice, 4)
	fmt.Println(meuSlice)

	meuSlice2 := meuSlice[1:3]
	fmt.Println(meuSlice2)

	meuSlice3 := meuSlice[:2]
	fmt.Println(meuSlice3)

	meuSlice4 := meuSlice[2:]
	fmt.Println(meuSlice4)

	meuSlice5 := meuSlice[:2]
	fmt.Println(meuSlice5)

	meuSlice6 := meuSlice[1:2]
	fmt.Println(meuSlice6)

	fmt.Println("\n=== Funções")
}
