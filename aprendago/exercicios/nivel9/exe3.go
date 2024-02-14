package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg_ sync.WaitGroup
var contador int

const quantidadeDeGoroutines = 10

func Exe3() {
	criaGoRoutines(quantidadeDeGoroutines)
	wg_.Wait()

	fmt.Println()
	fmt.Println("WaitGroup")
	fmt.Println("\tTotal de goroutines:\t", quantidadeDeGoroutines, "\n\tTotal do contador:\t", contador)
}

func criaGoRoutines(i int) {
	wg_.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg_.Done()
		}()
	}
}
