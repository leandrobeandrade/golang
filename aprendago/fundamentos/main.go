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
	fmt.Println("Sistemas numéricos")
	num := 31337
	fmt.Printf("Decimal: %d", num)
	fmt.Printf("\nBinário: %b", num)
	fmt.Printf("\nExadecimal: %#x\n", num)

	fmt.Println("================================")
	fmt.Println("Constantes")
	const con1 int = 100 // Tipada
	fmt.Println(con1)

	const con2 = 10 // Não tem o tipo atribuído até ser usada
	var var1 = 0
	var var2 int
	var fl float64

	var1 = con2 // Funciona (mesmo tipo)
	fmt.Println(var1)
	var2 = con2 // Funciona (mesmo tipo)
	fmt.Println(var2)
	fl = con2 // Funciona pq con2 só sabe o tipo na momento do uso
	fmt.Println(fl)
	// fl = var2 ERRO pq float não pode receber int

	const ( // Declaração múltipla
		um   = 1
		dois = 2
		tres = 3
	)
	fmt.Printf("%v %v %d\n", um, dois, tres)

	fmt.Println("================================")
	fmt.Println("Iota")
	const ( // omitindo valores
		quatro = iota
		_      = iota
		cinco  = iota
		seis   = iota
		_      = iota
	)
	fmt.Println(quatro, cinco, seis)

	const ( // preenchimento automático
		sete = iota
		_
		oito
		nove
		_
	)
	fmt.Println(sete, oito, nove)

	const ( // realizando operações
		dez = iota * 10
		_
		onze
		doze
		_
	)
	fmt.Println(dez, onze, doze)

	fmt.Println("================================")
	fmt.Println("Deslocamento de bits")
	d1 := 1
	d2 := d1 << 1
	fmt.Printf("%b\n", d1)
	fmt.Printf("%b\n", d2)

	d3 := 2
	d4 := d3 >> 1
	fmt.Printf("%b\n", d3)
	fmt.Printf("%b\n", d4)
}
