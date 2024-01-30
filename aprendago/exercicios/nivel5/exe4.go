package main

import "fmt"

func Exe4() {
	anon := struct {
		nome  string
		idade int
		map_  map[int]string
		sli   []int
	}{
		nome:  "Fulano",
		idade: 30,
		map_:  map[int]string{1: "UM", 2: "Dois", 3: "Três"},
		sli:   []int{10, 20, 30},
	}

	fmt.Println()
	fmt.Println(anon)
}
