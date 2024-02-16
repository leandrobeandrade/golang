package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func convergencia() {
	par := make(chan int)
	impar := make(chan int)
	conv := make(chan int)

	go pares(par, impar)
	go receber(par, impar, conv)

	fmt.Println()
	for v := range conv {
		fmt.Println("Valor recebido:", v)
	}

	fmt.Println("---------------------")

	canal := converge(trabalho("Montar Coisa"), trabalho("Embalar Coisa"))

	for i := 0; i < 20; i++ {
		fmt.Printf("Ciclo %v:\n  Mensagem: %v\n", i, <-canal)
	}
}

func pares(p, i chan int) {
	for n := 0; n <= 15; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func receber(p, i, c chan int) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(c)
}

func trabalho(texto string) chan string {
	canal := make(chan string)

	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Função %v diz: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(texto, canal)

	return canal
}

func converge(a, b chan string) chan string {
	novo := make(chan string)
	go func() {
		for v := range a {
			novo <- v
		}
	}()

	go func() {
		for v := range b {
			novo <- v
		}
	}()

	return novo
}
