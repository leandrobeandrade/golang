package main

import "fmt"

type pessoa_ struct {
	nome  string
	snome string
	idade int
}

type dentista struct {
	pessoa_
	espec   string
	salario float32
}

type arquiteto struct {
	pessoa_
	espec   string
	salario float32
}

// Polimorfismo - a mesma função [saudacao] trabalhando com classes diferentes [dentista - arquiteto]
func (x dentista) saudacao() {
	fmt.Println("Bom dia:", x.nome)
}

func (x arquiteto) saudacao() {
	fmt.Println("Bom dia:", x.nome)
}

// Interface com o método saudacao
type gente interface {
	saudacao()
}

// Função que implementa a interface
func ser_humano(g gente) {
	g.saudacao()
}

func Interfaces() {
	dentes := dentista{
		pessoa_: pessoa_{
			nome:  "Fulano",
			snome: "Silva",
			idade: 30,
		},
		espec:   "Extração",
		salario: 7.000,
	}

	predios := arquiteto{
		pessoa_: pessoa_{
			nome:  "Ciclano",
			snome: "Santos",
			idade: 35,
		},
		espec:   "Civil",
		salario: 16.000,
	}

	dentes.saudacao()
	predios.saudacao()

	fmt.Println()

	ser_humano(dentes)
	ser_humano(predios)
}
