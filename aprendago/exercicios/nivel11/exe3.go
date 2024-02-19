package main

import "fmt"

type erroEspecial struct {
	valor string
}

func (e erroEspecial) Erros() string {
	return fmt.Sprintf("Deu esse erro aqui: %v", e.valor)
}

func erroComoParametro(e erroEspecial) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.valor, "\nno método Error, eu tenho:", e)
}

func exe3() {
	arg := erroEspecial{"EMOJIS!!!!!!!!"}
	erroComoParametro(arg)
	fmt.Println()
}
