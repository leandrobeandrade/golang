package main

import "fmt"

type pessoa__ struct {
	nome  string
	snome string
	idade int
}

func (p pessoa__) mostra_pessoa() {
	fmt.Printf("\nNome completo: %v %v de idade: %d anos", p.nome, p.snome, p.idade)
}

func Exe4() {
	pessoa := pessoa__{"Fulano", "Santos", 32}
	pessoa.mostra_pessoa()
}
