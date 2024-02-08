package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Condicao() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines -", runtime.NumGoroutine())

	contador := 0
	totalGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			contador++
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("Goroutines <", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines >", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
