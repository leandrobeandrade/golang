package main

import (
	"fmt"
)

func Exe3() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
	fmt.Println()
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
