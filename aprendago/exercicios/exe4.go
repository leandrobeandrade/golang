package main

import "fmt"

type tipo int

var d tipo

func Exe4() {
	fmt.Printf("%v\n", d)
	fmt.Printf("%T\n", d)

	d = 42
	fmt.Println(d)
}
