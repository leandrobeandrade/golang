package main

import "fmt"

func Exe8() {
	rt := ret_func()
	rt()
}

func ret_func() func() {
	return func() {
		fmt.Println("Exercício 8, função que retorna função")
	}
}
