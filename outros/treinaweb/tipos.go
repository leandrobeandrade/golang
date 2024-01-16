package main

import (
	"fmt"
)

func tipos() {
	fmt.Println("\n===== TIPOS =====")
	fmt.Println("\n### Inteiros ###")

	// Vai de -128 a 127
	var inteiro8 int8 = -128
	fmt.Println("int8: ", inteiro8)

	// Vai de -32.768 a 32.767
	var inteiro16 int16 = 32767
	fmt.Println("int16: ", inteiro16)

	// Vai de -2.147.483.648 a 2.147.483.647
	var inteiro32 int32 = -2147483648
	fmt.Println("int32: ", inteiro32)

	// Vai de -9.223.372.036.854.775.808 a 9.223.372.036.854.775.807
	var inteiro64 int64 = 9223372036854775807
	fmt.Println("int64: ", inteiro64)

	fmt.Println("\n### Flutuantes ###")

	// Vai de -3.4e+38 a 3.4e+38
	var f32 float32 = 2.20
	fmt.Printf("float32: %.2f", f32)

	// Vai de -1.7e+308 a +1.7e+308 (tipo padrão)
	var f64 float64 = 459.689
	fmt.Print("\nfloat64: ", f64, "\n")

	fmt.Println("\n### String ###")
	txt1 := "Texto \"comum\""
	fmt.Println(txt1)

	// Não interpretam caracteres especiais, exemplo \ vai aparecer no texto, bom pra quebrar linhas no texto
	txt2 := `Texto com "crase"`
	fmt.Println(txt2)

	// Carateres UTF-8
	txt3 := "Japonês, 世界"
	fmt.Println(txt3)

	// Mesmo que tipo rune
	txt4 := 'A'
	fmt.Println(txt4)

	fmt.Println("\n### Rune ###")
	// Armazena códigos que representam caracteres Unicode
	var rune_ rune = 'A'
	fmt.Println("rune: ", rune_)

	// Declarando e inicializando com um caractere Unicode
	emoji := "😀"
	fmt.Println("rune: ", emoji)

	fmt.Println("\n### Boolean ###")
	bol := true
	fmt.Println(bol)
	fmt.Println(!bol)

	fmt.Println("\n### Slice ###")
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	fmt.Println("\n### Map ###")
	map_ := map[rune]int{'A': 1, 'B': 2, 'C': 3}
	vl := 0

	for _, v := range map_ {
		vl += v
	}
	fmt.Println(vl)
}
