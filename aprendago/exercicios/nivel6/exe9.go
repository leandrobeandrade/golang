package main

import "fmt"

func Exe9() {
	y := dois
	y(um)
}

func um() {
	fmt.Println("exercício 9, funcão como callback")
}

func dois(x func()) {
	x()
}
