package main

import "fmt"

func Exe10() {
	maps := map[string][]string{
		"Silva_Fulano ":    {"Futebol", "Hóquei"},
		"Santos_Ciclano":   {"Volei", "Arco e flecha"},
		"Pereira_Beltrano": {"Natação", "Atletismo"},
	}

	maps["Lima_Belciclano"] = []string{"Basquete", "Hipismo", "Balonismo"}

	delete(maps, "Pereira_Beltrano")

	fmt.Println()
	for k, v := range maps {
		fmt.Println(k)

		for i, vl := range v {
			fmt.Println("\t", i, vl)
		}
	}
}
