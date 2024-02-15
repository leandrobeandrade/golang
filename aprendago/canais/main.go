package main

import "fmt"

func main() {
	canal := make(chan int)

	// canal <- 42 // ERRO

	go func() {
		canal <- 10
	}()

	fmt.Println("Valor canal:", <-canal)

	go send(canal)
	receive(canal)

	rangeClose()
	buffer()
	canalGeralParaEspecifico()
	select_()
	select__()
	select___()
	convergencia()
	divergencia()
	context_()
}

func send(s chan<- int) {
	s <- 20
}

func receive(r <-chan int) {
	fmt.Println("\nValor recebido do canal alterado:", <-r)
}
