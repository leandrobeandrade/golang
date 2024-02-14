package main

import "fmt"

func select_() {
	a := make(chan int)
	b := make(chan int)
	c := 10

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(c / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(c / 2)

	fmt.Println()

	for i := 0; i < c; i++ {
		select {
		case v := <-a:
			fmt.Println("Canal A:", v)
		case v := <-b:
			fmt.Println("Canal B:", v)
		}
	}
}

func select__() {
	canal := make(chan int)
	quit := make(chan int)

	fmt.Println()

	go func(c chan int) {
		for i := 0; i < 20; i++ {
			fmt.Println("Recebido:", <-c)
		}
		quit <- 0
	}(canal)

	func(c chan int) {
		qualquerCoisa := 10

		for {
			select {
			case c <- qualquerCoisa:
				qualquerCoisa++
			case <-quit:
				return
			}
		}
	}(canal)
}

func select___() {
	par := make(chan int)
	impar := make(chan int)
	quit := make(chan bool)

	fmt.Println()
	go envia(par, impar, quit)
	recebe(par, impar, quit)
}

func envia(par_, impar_ chan int, quit_ chan bool) {
	for i := 0; i <= 30; i++ {
		if i%2 == 0 {
			par_ <- i
		} else {
			impar_ <- i
		}
	}
	quit_ <- true
	close(par_)
	close(impar_)
}

func recebe(par_, impar_ chan int, quit_ chan bool) {
	for {
		select {
		case v := <-par_:
			fmt.Println("O número:", v, "é par!")
		case v := <-impar_:
			fmt.Println("O número:", v, "é ímpar!")
		case <-quit_:
			return
		}
	}
}
