package main

import "fmt"

func Exe1() {
	fmt.Println(ret_int())
	fmt.Println(ret_string_int())
}

func ret_int() int {
	return 10
}

func ret_string_int() (string, int) {
	x := 10
	return "Valor:", x
}
