package main

import "fmt"

func Exe2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()

	fmt.Println()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
	fmt.Println()
}
