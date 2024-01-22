package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	c := "😁"
	fmt.Println(a, b, c)

	d := []byte(a) // 1 byte
	e := []byte(b) // 2 bytes
	f := []byte(c) // 3 bytes
	fmt.Println(d, e, f)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	var i uint16 = 65535 // n° máximo suportado
	fmt.Println(i)
	i++ // overflow -> começa a contagem
	fmt.Println(i)
	i++
	fmt.Println(i)

	s := `String 
			em múltipla linhas`
	fmt.Printf("Tipo: %T\t Valor: %v", s, s)
	fmt.Print("\n")

	s2 := "ascii éøâ 香"
	fmt.Printf("\n%v", s2)
	fmt.Println()

	for _, v := range s2 { // for com range devolve caracteres
		fmt.Printf("%b - %T - %#U %#x\n", v, v, v, v)
	}

	fmt.Println()

	for i := 0; i < len(s2); i++ { // for loop devolve bits
		fmt.Printf("%b - %T - %#U %#x\n", s2[i], s2[i], s2[i], s2[i])
	}

	fmt.Println("================================")

}
