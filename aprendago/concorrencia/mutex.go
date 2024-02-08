package main

import (
	"fmt"
	"runtime"
	"sync"
)

func Mutex() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines -", runtime.NumGoroutine())

	contador := 0
	totalGoRoutines := 10

	var wg sync.WaitGroup
	wg.Add(totalGoRoutines)

	var mu sync.Mutex

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			mu.Lock()
			contador++
			runtime.Gosched()
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines <", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines >", runtime.NumGoroutine())
	fmt.Println("Valor final", contador)
}
