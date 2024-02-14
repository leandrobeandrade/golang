package main

import (
	"fmt"
	"runtime"
)

func Exe6() {
	fmt.Println()
	fmt.Println("Seu OS é:\t", runtime.GOOS)
	fmt.Println("Seu ARCH é:\t", runtime.GOARCH)
}
