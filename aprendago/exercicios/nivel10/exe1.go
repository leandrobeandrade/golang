package main

import "fmt"

func Exe1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)

	c_ := make(chan int, 1) // Buffer
	c_ <- 10

	fmt.Println(<-c_)
}
