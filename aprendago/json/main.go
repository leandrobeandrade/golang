package main

import (
	"encoding/json"
	"fmt"
)

type pessoaJson struct {
	Nome          string
	Snome         string
	Idade         int
	Profis        string
	ContaBancaria float64
}

func main() {
	jamesBond := pessoaJson{
		Nome:          "James",
		Snome:         "Bond",
		Idade:         40,
		Profis:        "Agente secreto",
		ContaBancaria: 1000000.50,
	}

	darthVader := pessoaJson{"Dart", "Vader", 50, "Vilão intergaláctico", 5000000000.83}

	fmt.Println(jamesBond)
	fmt.Println(darthVader)
	fmt.Println()

	j, err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))

	d, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))
}
