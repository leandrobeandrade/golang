package main

import "fmt"

var a int = 42
var b string = "James Bond"
var c bool = true

func Exe3() {
	s := fmt.Sprintf("%v\n%v\n%v\n", a, b, c)
	fmt.Println(s)
}
