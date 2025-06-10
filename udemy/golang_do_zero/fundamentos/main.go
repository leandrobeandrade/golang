package main

import (
	"errors"
	"fmt"
	"udemy/golang_do_zero/fundamentos/funcs"
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
	var numero int = 1000000000000000000 // int usa a arquitetura do sistema
	fmt.Println(numero)

	var numero1 int16 = -10000 // int16 é um tipo de dado de 16 bits
	fmt.Println(numero1)

	var numero2 int32 = 1000000000 // int32 é um tipo de dado de 32 bits
	fmt.Println(numero2)

	var numero3 int64 = -1000000000000000000 // int64 é um tipo de dado de 64 bits
	fmt.Println(numero3)

	var numero4 uint8 = 100 // uint8 é um tipo de dado de 8 bits e positivo apenas
	fmt.Println(numero4)

	var numero5 rune = 10000 // rune é um tipo de dado de 32 bits equivalente a int32
	fmt.Println(numero5)

	var numero6 byte = 100 // byte é um tipo de dado de 8 bits equivalente a uint8
	fmt.Println(numero6)

	var numeroReal float32 = 123.456
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 1234567890000.456
	fmt.Println(numeroReal2)

	numeroReal3 := 123.456 // float implicito utiliza a arquitetura do sistema
	fmt.Println(numeroReal3)

	fmt.Println("\n=== Strings")
	var str string = "String"
	fmt.Println(str)

	var str2 string = `
		String
		com
		quebra
		de
		linha
	`
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	fmt.Println("\n=== Booleanos")
	var booleano bool = true
	fmt.Println(booleano)

	fmt.Println("\n=== Erros")
	var erro error = nil
	fmt.Println(erro)

	var err error = errors.New("erro interno")
	fmt.Println(err)

	fmt.Println("\n=== Arrays")
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
	funcs.AllFuncs()
}
