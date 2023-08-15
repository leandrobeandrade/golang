package main

import ("fmt")

type Person struct {
	name  string
	age   int
	email string
}

func main() {
	person := Person{
	  name:  "Fulano",
		age:   35,
		email: "fulano@gmail.com",
	}

	fmt.Println("\n### struct ###")
	fmt.Println("A pessoa tem o nome ", person.name, " a idade ", person.age, " anos e o email: ", person.email)
}