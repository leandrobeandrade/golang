package main

import (
	"encoding/json"
	"fmt"
)

func Exe2() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	json_, err := json.Marshal(users)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(json_))
}
