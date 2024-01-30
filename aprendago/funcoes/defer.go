package main

import "fmt"

func Defer() {
	fmt.Println()
	defer fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}
