package main

import "fmt"

func Exe6() {
	const (
		a = 2025 + iota
		b
		c
		d
	)
	fmt.Printf("\n%d %d %d %d", a, b, c, d)
}
