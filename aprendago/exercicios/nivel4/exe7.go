package main

import "fmt"

func Exe7() {
	sli := [][]string{
		{"Fulano", "Silva", "Esporte"},
		{"Ciclano", "Santos", "Filmes"},
		{"Beltrano", "Pereira", "Games"},
	}

	fmt.Println()
	for _, v := range sli {
		fmt.Println(v)
	}

	for _, v := range sli {
		fmt.Println(v[0])
		for _, v2 := range v {
			fmt.Println("\t", v2)
		}
	}
}
