package main

import "fmt"

type pessoa___ struct {
	nome  string
	snome string
	idade int
}

func Exe2() {
	pess := pessoa___{"Fulano", "Silva", 30}

	fmt.Print("\n", pess, "\n")

	mudeMe(&pess)

	fmt.Println(pess)
}

func mudeMe(p *pessoa___) {
	(*p).nome = "Beltrano" // as 2 maneiras funcionam
	p.snome = "Santos"
}
