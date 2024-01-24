package main

import (
	"fmt"
	"time"
)

func Exe2() {
	for i := 1984; i <= time.Now().Year(); i++ {
		fmt.Println(i)
	}
	fmt.Println()
}
