package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("exem.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, err := io.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
