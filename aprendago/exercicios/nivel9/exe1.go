package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Exe1() {
	wg.Add(2)

	go um()
	go dois()

	wg.Wait()
}

func um() {
	fmt.Println("Função um disparada!")
	wg.Done()
}

func dois() {
	fmt.Println("Função dois disparada!")
	wg.Done()
}

// Outro exemplo
/* func novasgoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("Eu sou a goroutine número:", i+1)
			wg.Done()
		}(x)
	}
}

func main() {
	novasgoroutines()
	wg.Wait()
} */
