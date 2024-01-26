package main

import "fmt"

func Exe8() {
	maps := map[string][]string{
		"Silva_Fulano ":    {"Futebol", "Hóquei"},
		"Santos_Ciclano":   {"Volei", "Arco e flecha"},
		"Pereira_Beltrano": {"Natação", "Atletismo"},
	}

	fmt.Println()
	for k, v := range maps {
		fmt.Println(k)

		for i, vl := range v {
			fmt.Println("\t", i, vl)
		}
	}
}
