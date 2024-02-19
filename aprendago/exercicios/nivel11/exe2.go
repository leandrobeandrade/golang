package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person_ struct {
	First   string
	Last    string
	Sayings []string
}

func exe2() {
	p1 := person_{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln("Erro.....")
	}

	fmt.Println(string(bs))
	fmt.Println()
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("erro Aqui :)")
	}

	return bs, err
}
