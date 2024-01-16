package ponteiros

import "fmt"

func Ponteiros() {
	a := 10
	fmt.Println("Valor de a:", a)
	fmt.Println("Endereço de a:", &a)

	var ponteiro *int = &a
	fmt.Println("Copiando endereço de a:", ponteiro)
	fmt.Println("Endereço de ponteiro:", &ponteiro)
	fmt.Println("Valor de ponteiro:", *ponteiro)

	// mudando o valor pelo ponteiro
	*ponteiro = 50
	fmt.Println("Novo valor de ponteiro:", *ponteiro)
	fmt.Println("Novo valor de a:", a)
}

// Abc muda o valor passado por parametro através de um ponteiro
func Abc(a *int) int {
	*a = 200
	return *a
}

type Carro struct {
	Nome string
}

// Mudou() muda o valor pelo endereço de memoria de Carro, ou seja em todas as chamadas de Carro terá o novo valor
func (c *Carro) Mudou() {
	c.Nome = "BMW"
	fmt.Println(c.Nome, "mudou!")
}
