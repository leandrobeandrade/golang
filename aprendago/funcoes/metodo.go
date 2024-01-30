package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) pertence_ao_objeto_pessoa() {
	fmt.Println()
	fmt.Println(p.nome)
	fmt.Println()
}

func Metodo() {
	metodo_ := pessoa{nome: "Fulano", idade: 35}
	metodo_.pertence_ao_objeto_pessoa()
}
