package main

import "fmt"

func rangeClose() {
	canal := make(chan int)

	go meuloop(10, canal)
	prints(canal)
}

func meuloop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s) // encerra o canal
}

func prints(r <-chan int) {
	fmt.Println()
	for v := range r {
		fmt.Println("Recebido do canal:", v)
	}
}
