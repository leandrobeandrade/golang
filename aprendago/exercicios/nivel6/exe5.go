package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float32
}

type circulo struct {
	raio float32
}

type figura interface {
	area()
}

func info(f figura) {
	f.area()
}

func (q quadrado) area() {
	resultado := q.lado * q.lado
	fmt.Println("Área do quadrado:", resultado)
}

func (c circulo) area() {
	resultado := math.Pi * c.raio * c.raio
	fmt.Println("Área do círculo:", resultado)
}

func Exe5() {
	x := quadrado{
		lado: 15.0,
	}

	y := circulo{
		raio: 5.25,
	}

	fmt.Print("\n\n")
	x.area()
	y.area()
	info(x)
	info(y)
}
