package main

import "fmt"

func Exe7() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("quer jogar futebol, vai pro Brasil")
	case "starcraft":
		fmt.Println("quer jogar starcraft, vai pra Coréia")
	case "espeleísmo":
		fmt.Println("quer fazer essa coisa estranha, vai pro psiquiatra")
	}
}
