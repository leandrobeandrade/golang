package main

import "fmt"

func buffer() {
	canal := make(chan int, 1) // buffer - armazena uma execução

	canal <- 30
	// canal2 <- 40 // Erro - o Buffer só armazena 1 execução
	fmt.Println("\nValor do canal com buffer:", <-canal)
}
