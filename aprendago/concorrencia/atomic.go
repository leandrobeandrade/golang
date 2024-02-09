package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Função atomica -> executa do começo ao fim não em partes
func Atomic() {
	fmt.Println("CPUs:", runtime.NumCPU())

	var contador int64 = 0
	totalGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Valor final", contador)
}
