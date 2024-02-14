package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p *pessoa) falar() {
	fmt.Println(p.nome, "diz oi!")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(hum humanos) {
	hum.falar()
}

func Exe2() {
	pessoa1 := pessoa{
		nome:  "Genghis Khan",
		idade: 1000,
	}

	pessoa2 := pessoa{
		nome:  "Sun Tzu",
		idade: 2500,
	}

	fmt.Println()
	pessoa1.falar()
	(&pessoa1).falar() // Também funciona - mais recomendável
	dizerAlgumaCoisa(&pessoa2)
}
