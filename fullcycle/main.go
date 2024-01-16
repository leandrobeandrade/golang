package main

import (
	"fmt"
	"fullcycle/context"
	funcoes "fullcycle/function"
	generics "fullcycle/generic"
	ponteiros "fullcycle/point"
	rotinas "fullcycle/routine"
	api "fullcycle/service"
	"fullcycle/structs"
	"time"
)

func main() {
	cond := true

	if !cond {
		// Funções
		fmt.Println(funcoes.Soma1(1, 1))
		fmt.Println(funcoes.Soma2(2, 2))
		soma := funcoes.Soma3(3, 3)
		fmt.Println(soma)
		fmt.Println(funcoes.Soma4(4 + 4))
		funcoes.Soma5()
		num_ := funcoes.Nums{
			Num1: 10,
			Num2: 10,
		}
		num_.Soma6()

		// Ponteiros
		ponteiros.Ponteiros()
		exemp := 10
		fmt.Println(ponteiros.Abc(&exemp))

		carro := ponteiros.Carro{
			Nome: "Ka",
		}
		fmt.Println(carro.Nome)
		carro.Mudou()
		fmt.Println(carro.Nome)

		// Structs
		structs.Cliente1()
		structs.Cliente2()
		structs.Cliente3()
		structs.Cliente4()

		// Go Routine
		go rotinas.Contador("sem rotina")
		go rotinas.Contador("com rotina")
		time.Sleep(time.Second * 5)
		fmt.Println("...FIM...")

		rotinas.Chanel1()
		rotinas.Select()

		msg := make(chan int)
		go rotinas.Worker(1, msg) // cada Worker compartilham a mesma msg
		go rotinas.Worker(2, msg)
		for i := 0; i < 10; i++ {
			msg <- i // preenche msg com o valor de i
		}

		// Generics
		fmt.Println("Soma com generics INT:", generics.Soma1(map[string]int64{"a": 1, "b": 2, "c": 3}))
		fmt.Println("Soma com generics FLOAT:", generics.Soma1(map[string]float64{"a": 1.1, "b": 2.2, "c": 3}))

		fmt.Println("Soma com generics INT tipada:", generics.Soma2(map[string]int64{"a": 1, "b": 2, "c": 3}))
		fmt.Println("Soma com generics FLOAT tipada:", generics.Soma2(map[string]float64{"a": 1.1, "b": 2.2, "c": 3}))
		fmt.Println("Soma com generics STRING tipada:", generics.Soma2(map[string]rune{"a": 'A', "b": 'B', "c": 'C'}))

		var x, y, z generics.NumInt
		x = 1
		y = 2
		z = 3
		fmt.Println("Soma com generics INT com 2 types:", generics.Soma3(map[string]generics.NumInt{"a": x, "b": y, "c": z}))
		fmt.Println("Soma com generics INT com 2 types:", generics.Soma3(map[string]generics.NumInt{"a": 'X', "b": 'Y', "c": 'Z'}))

		fmt.Println("Soma com generics INT por parametros:", generics.Soma4(10.0, 10.0))

		fmt.Println("Comparando com interface Comparable:", generics.Compara(10.0, 12.1))
		fmt.Println("Comparando com interface Ordered:", generics.Max[int64](60, 50))
		fmt.Println("Concatenando:", generics.Concat([]generics.MyString{"A", "B"}))

		// Http
		api.Read()

		t := f()
		fmt.Println(t)
	} else {
		context.Exec()
	}

}

func f() []int {
	s := make([]int, 100)
	return s
}
