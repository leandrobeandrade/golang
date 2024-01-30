package main

import "fmt"

type pessoa_ struct {
	Nome      string
	Sobrenome string
	Sabores   []string
}

func Exe2() {
	map1 := make(map[string]pessoa_)

	map1["Silva"] = pessoa_{
		Nome:      "Fulano",
		Sobrenome: "Silva",
		Sabores:   []string{"Chocolate", "Flocos", "Abacaxi"},
	}
	map1["Santos"] = pessoa_{"Beltrano", "Santos", []string{"Coco", "Morango", "Uva", "Leite"}}
	fmt.Println()

	for _, v := range map1 {
		fmt.Println(v)
	}

	// OU

	map2 := map[string]pessoa_{
		"Silva": {
			Nome:      "Fulano",
			Sobrenome: "Silva",
			Sabores:   []string{"Chocolate", "Flocos", "Abacaxi"},
		},
		"Santos": {"Beltrano", "Santos", []string{"Coco", "Morango", "Uva", "Leite"}},
	}
	fmt.Println()

	for _, v := range map2 {
		fmt.Println(v)
	}

	fmt.Println()
	for _, v := range map1 {
		fmt.Println("Nome:", v.Nome, "sorvetes favoritos:")
		for _, sab := range v.Sabores {
			fmt.Println("-", sab)
		}
	}
	fmt.Print("\n")
}
