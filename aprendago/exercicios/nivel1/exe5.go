package main

import "fmt"

var e int

func Exe5() {
	fmt.Printf("\n%v\n", d)
	fmt.Printf("%T\n", d)

	d = 42
	fmt.Println(d)

	e = int(d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
}
