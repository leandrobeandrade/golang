package rotinas

import (
	"fmt"
	"time"
)

func Contador(tipo string) {
	for i := 0; i < 5; i++ {
		fmt.Println(tipo, i)
		time.Sleep(time.Second)
	}
}

// Thread 1
func Chanel1() {
	// Thread 1 comunica com a Thread 2
	hello := make(chan string)

	// Thread 2
	go func() {
		hello <- "Hello World"
	}()

	resultado := <-hello

	fmt.Println(resultado)
}

func Select() {
	hello := make(chan string)

	go func() {
		hello <- "hello world"
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("Default")
	}

}

func Worker(workerId int, msg chan int) {
	for res := range msg {
		fmt.Println("Worker:", workerId, "Msg:", res)
		time.Sleep(time.Second)
	}
}

// https://www.youtube.com/watch?v=B4NL0rMvXMg&list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg&index=8&t=387s&ab_channel=FullCycle
