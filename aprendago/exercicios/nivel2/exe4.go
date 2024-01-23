package main

import "fmt"

func Exe4() {
	a := 10
	fmt.Printf("\n%d\t %b\t %#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t %b\t %#x\n", b, b, b)
}
