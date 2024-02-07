package main

import (
	"fmt"
	"sort"
)

type user__ struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func Exe5() {
	u1 := user__{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user__{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user__{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user__{u1, u2, u3}

	fmt.Println()
	fmt.Println(users)

	sort.Slice(users, func(i, j int) bool {
		return users[i].Age < users[j].Age || (users[i].Age == users[j].Age && users[i].Last < users[j].Last)
	})

	for i, u := range users {
		fmt.Printf("%d. %s %s, %d:\n", i+1, u.First, u.Last, u.Age)

		for j, s := range u.Sayings {
			fmt.Printf("\t%d. %s\n", j+1, s)
		}

		fmt.Println()
	}

	fmt.Println("-- Outra Maneira --")
	printUsers(users)
}

func printUsers(users []user__) {
	for i, user := range users {
		fmt.Printf("%v\t Name: %v %v Age: %v\n \t Sayings:\n", i+1, user.First, user.Last, user.Age)
		for _, say := range user.Sayings {
			fmt.Printf("\t\t%v\n", say)
		}
		fmt.Println()
	}
}
