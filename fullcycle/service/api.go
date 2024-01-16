package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Read() {
	req, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	defer req.Body.Close()
	resp, _ := io.ReadAll(req.Body)

	fmt.Println(string(resp))
}
