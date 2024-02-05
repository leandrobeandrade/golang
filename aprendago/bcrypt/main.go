package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	senha := "01012000"
	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Hash criado a partir da senha:", string(sb))

	comp_certa := bcrypt.CompareHashAndPassword(sb, []byte(senha))
	fmt.Println("Validação sem erro:", comp_certa)

	senha_errada := "01012001"
	comp_error := bcrypt.CompareHashAndPassword(sb, []byte(senha_errada))
	fmt.Println("Validação com erro", comp_error)
}
