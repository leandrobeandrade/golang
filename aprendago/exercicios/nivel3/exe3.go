package main

import (
	"fmt"
	"time"
)

func Exe3() {
	i := 1984

	for {
		if i > time.Now().Year() {
			break
		}
		fmt.Println(i)
		i++
	}
}
