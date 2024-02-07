package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user_ struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func Exe3() {
	u1 := user_{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user_{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user_{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user_{u1, u2, u3}

	encode := json.NewEncoder(os.Stdout)
	encode.Encode(users)

	fmt.Println()
	fmt.Println("___ Forma 1:", users)

	// Outra forma - forma direta
	fmt.Print("___ Forma 2:")
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}
}
