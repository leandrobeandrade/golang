package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
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

	j, err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n=== Marshal ===")
	fmt.Println(string(j))

	d, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(d))

	sb := []byte(`{"Nome": "James", "Snome": "Bond", "Idade": 40, "Profis": "Agente secreto", "ContaBancaria": 1000000.5}`)
	var dados pessoaJson

	err = json.Unmarshal(sb, &dados)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\n=== Unmarshal ===")
	fmt.Println(dados)

	encoder := json.NewEncoder(os.Stdout)

	fmt.Println("\n=== Encoder ===")
	encoder.Encode(jamesBond) // Não precisa passar para uma variável como o Marshal

	fmt.Println("\n=== Sort ===")
	ss := []string{"abóbora", "uva", "maçã", "laranja", "jaca"}

	sort.Strings(ss)
	fmt.Println(ss)

	si := []int{10, 35, 01, 50}

	sort.Ints(si)
	fmt.Println(si)

	fmt.Println("\n=== Sort Customizado")

	carros := []carro{
		{nome: "Chevete", potencia: 50, consumo: 5},
		{nome: "Porsche", potencia: 300, consumo: 15},
		{nome: "Fusca", potencia: 20, consumo: 30},
	}

	fmt.Println("Dado Inicial:\t\t", carros)
	sort.Sort(ordenaPorPotencia(carros))

	fmt.Println("Ordena por Potência:\t", carros)
	sort.Sort(ordenaPorConsumo(carros))

	fmt.Println("Ordena por Consumo:\t", carros)
	sort.Sort(ordenaPorLucro(carros))

	fmt.Println("Ordena por Lucro:\t", carros)
	sort.Sort(ordenaPorLucro(carros))
}
