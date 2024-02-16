package main

import (
	"fmt"
)

func Exe4() {
	q := make(chan int)
	c := gen_(q)

	receive_(c, q)

	fmt.Println("about to exit")
	fmt.Println()
}

func gen_(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 15; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()

	return c
}

func receive_(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Valores do canal:", v)
		case <-q:
			return
		}
	}
}
