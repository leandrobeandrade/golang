package main

import "fmt"

func Closure() {
	a := func_closure()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := func_closure()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
}

func func_closure() func() int {
	x := 0

	return func() int {
		x++
		return x
	}
}
