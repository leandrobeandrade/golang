package main

import "fmt"

type client struct {
	nome    string
	snome   string
	fumante bool
}

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	client1 := client{
		nome:    "Fulano",
		snome:   "Silva",
		fumante: false,
	}

	client2 := client{"Beltrano", "Sousa", false} // atribuição direta de valores, ocultando as propriedades

	fmt.Println(client1)
	fmt.Println(client2)
	fmt.Println()

	pessoa1 := pessoa{
		nome:  "Ciclano",
		idade: 30,
	}

	pessoa2 := profissional{
		pessoa:  pessoa1,
		titulo:  "Lord",
		salario: 30000,
	}

	pessoa3 := profissional{
		pessoa: pessoa{
			nome:  "Belciclano",
			idade: 35,
		},
		titulo:  "Sir",
		salario: 55000,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	fmt.Println(pessoa2.pessoa)
	fmt.Println()
	fmt.Println(pessoa3)
	fmt.Println(pessoa3.pessoa.nome)
	fmt.Println(pessoa3.nome) // consegue acessar elementos da struct interna se as propriedades não forem iguais
	fmt.Println(pessoa3.titulo)

	pessoa4 := struct {
		nome  string
		idade int
	}{
		nome:  "John Doe",
		idade: 55,
	}

	fmt.Println()
	fmt.Println(pessoa4)
}
