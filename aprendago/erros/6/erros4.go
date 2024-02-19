package main

import (
	"fmt"
	"log"
)

func erros4() {
	_, err := sqrt4(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt4(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	return 42, nil
}
