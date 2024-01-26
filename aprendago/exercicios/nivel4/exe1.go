package main

import "fmt"

func Exe1() {
	arr := [5]int{1, 2, 3, 4, 5}

	for i, v := range arr {
		fmt.Printf("\n%T %v %v", v, i, v)
	}
	fmt.Println()
	fmt.Println()
}
