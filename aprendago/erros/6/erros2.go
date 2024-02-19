package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func erros2() {
	fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt2(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt2(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorgateMath
	}
	return 42, nil
}

// see use of errors.New in standard library:
// http://golang.org/src/pkg/bufio/bufio.go
// http://golang.org/src/pkg/io/io.go
