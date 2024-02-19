package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err happened FMT:", err)
		log.Println("\nerr happened LOG:", err)
		log.Fatalln("\nerr happened FATAL", err)
		panic(err)
	}
}
