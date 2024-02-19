package main

import (
	"fmt"
	"log"
)

func erros3() {
	_, err := sqrt3(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt3(f float64) (float64, error) {
	if f < 0 {
		return 0, fmt.Errorf("norgate math again: square root of negative number: %v", f)
	}
	return 42, nil
}
