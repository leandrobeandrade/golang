package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	defer fmt.Println("Rodou a função")

	r := strings.NewReader("James Bond")
	io.Copy(f, r)
}
