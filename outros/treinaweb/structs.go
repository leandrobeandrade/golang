package main

import ("fmt")

type Pessoa struct {
  nome  string
  idade int
  casado bool
}

func main() {
  pessoa := Pessoa {
    nome:  "Fulano",
    idade:   35,
    casado: true,
  }

  fmt.Println("A pessoa tem o nome ", pessoa.nome, " a idade de ", pessoa.idade, " anos e casado: ", pessoa.casado)
}