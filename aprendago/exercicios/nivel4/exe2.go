package main

import "fmt"

func Exe2() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range sli {
		fmt.Print(v, "\t")
	}
}
