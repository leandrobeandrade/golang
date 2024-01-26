package main

import "fmt"

func Exe3() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Print("\n\n")
	fmt.Println(sli[:3])
	fmt.Println(sli[4:])
	fmt.Println(sli[1:7])
	fmt.Println(sli[2:9])
	fmt.Println(sli[2 : len(sli)-1])
}
