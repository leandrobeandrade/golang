package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg__ sync.WaitGroup
var contador_ int
var mu sync.Mutex

const quantidadeDeGoroutines_ = 10

func Exe4() {
	criaGoRoutines_(quantidadeDeGoroutines_)
	wg__.Wait()

	fmt.Println()
	fmt.Println("Mutex")
	fmt.Println("\tTotal de goroutines:\t", quantidadeDeGoroutines_, "\n\tTotal do contador:\t", contador_)
}

func criaGoRoutines_(i int) {
	wg__.Add(i)

	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador_
			runtime.Gosched()
			v++
			contador_ = v
			mu.Unlock()
			wg__.Done()
		}()
	}
}
