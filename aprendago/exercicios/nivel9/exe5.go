package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func Exe5() {
	var wg sync.WaitGroup
	var contador int64 = 0
	totalGoRoutines := 20

	wg.Add(totalGoRoutines)

	fmt.Println("Atomic")

	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			v := atomic.LoadInt64(&contador)
			runtime.Gosched()
			fmt.Println("\t", v)
			wg.Done()
		}()
	}
}
