package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("\nRecovered in f", r)
		}
	}()
	fmt.Println("\nCalling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("\tDefer in g", i)
	fmt.Println("\tPrinting in g", i)
	g(i + 1)
}
