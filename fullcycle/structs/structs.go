package structs

import (
	"encoding/json"
	"fmt"
	"log"
)

type Cliente struct {
	Nome  string
	Email string
	CPF   int
}

type ClienteInternacional struct {
	Cliente
	Pais string `json:"País"` // tag
}

func Cliente1() {
	cliente := Cliente{
		Nome:  "Fulano",
		Email: "ful@email.com",
		CPF:   12345678900,
	}

	fmt.Println(cliente)
	cliente.ExibeCliente()
}

func Cliente2() {
	cliente := Cliente{"Beltrano", "bel@email.com", 98765432111}

	fmt.Printf("Nome: %s, E-mail: %s, CPF: %d \n", cliente.Nome, cliente.Email, cliente.CPF)
	cliente.ExibeCliente()
}

func Cliente3() {
	cliente := ClienteInternacional{
		Cliente: Cliente{
			Nome:  "Ciclano",
			Email: "cic@email.com",
			CPF:   123212123,
		},
		Pais: "EUA",
	}

	fmt.Printf("Nome: %s, E-mail: %s, CPF: %d, País: %s\n", cliente.Nome, cliente.Email, cliente.CPF, cliente.Pais)
	cliente.ExibeCliente()
	ClienteJson(cliente.Cliente)
}

// ExibeCliente -> método pertencente a struct Cliente
func (c Cliente) ExibeCliente() {
	fmt.Println("Exibe nome do cliente:", c.Nome)
}

// Marshal converte struct para json
func ClienteJson(cliente Cliente) {
	clienteJson, err := json.Marshal(cliente)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clienteJson))
}

// Unmarshal converte json em struct
func Cliente4() {
	jsonCliente := `{"Nome": "BelCiclano", "Email": "bec@email.com", "CPF": 91827364500, "País": "Canadá"}`
	cliente := ClienteInternacional{}

	json.Unmarshal([]byte(jsonCliente), &cliente)
	fmt.Println(cliente)
}
