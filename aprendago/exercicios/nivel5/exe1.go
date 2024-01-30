package main

import "fmt"

type pessoa struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func Exe1() {
	pessoa1 := pessoa{
		Nome:      "Fulano",
		Sobrenome: "Silva",
		Sabores:   []string{"Chocolate", "Flocos", "Abacaxi"},
	}
	fmt.Println(pessoa1)

	pessoa2 := pessoa{"Beltrano", "Santos", []string{"Coco", "Morango", "Uva", "Leite"}}
	fmt.Print(pessoa2, "\n\n")

	for _, v := range pessoa1.Sabores {
		fmt.Println(v)
	}

	fmt.Print("\n")
	for _, v := range pessoa2.Sabores {
		fmt.Println(v)
	}
}
