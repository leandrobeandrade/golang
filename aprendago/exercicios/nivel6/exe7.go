package main

import "fmt"

func Exe7() {
	f := funcao_atribuida
	f()
}

func funcao_atribuida() {
	fmt.Println("exercício 7, função atribuída")
}
