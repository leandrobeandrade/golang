package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracao bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func Exe3() {
	caminhonete_ := caminhonete{veiculo{6, "Vermelha"}, true}
	sedan_ := sedan{veiculo{4, "verde"}, false}

	fmt.Println(caminhonete_)
	fmt.Println(caminhonete_.portas)
	fmt.Println(sedan_)
	fmt.Println(sedan_.veiculo.cor)
}
