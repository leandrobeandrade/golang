package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// Roda sequencialmente as 2 funções
	/* um()
	dois() */

	// fmt.Println()

	// Roda apenas a função dois pq chega no final da main e para a execução
	// Função um fica aguardando e não roda
	/* go um()
	dois() */

	// fmt.Println()

	// Roda as 2 funções sequencialmente
	/* wg.Add(1)
	go um()
	dois()
	wg.Wait() */

	// Roda as 2 funções concorrentemente
	wg.Add(2)
	go um()
	go dois()
	wg.Wait()

	fmt.Println()
	Condicao()
	fmt.Println()
	Mutex()
}

func um() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Um:", i)
		time.Sleep(1000)
	}
	wg.Done()
}

func dois() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Dois:", i)
		time.Sleep(1000)
	}
	wg.Done()
}
