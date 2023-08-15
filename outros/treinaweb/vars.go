package main

import "fmt"

func main() {

  // declaração com inferência do tipo
  var idade = 10
  fmt.Println("Olá Mundo GO >>>>", idade)

  // declaração e atribuição direta
	var idade1 int = 18
  fmt.Println("Olá Mundo GO >>>>", idade1)

  // declaração tipada com atribuição posterior
	var idade2 int
	idade2 = 30

	fmt.Println("Olá Mundo GO >>>>", idade2)

  // operador gopher só funciona em funções e com inferência do tipo
  idade3 := 15
  fmt.Println("Olá Mundo GO >>>>", idade3)

  // constantes devem ser inicializadas e não precisam ser tipadas
  const PI float64 = 3.141592653589793
  fmt.Println("const", PI)

  bol := true
  fmt.Println(bol != true)

  x := 5
  y := 8

  fmt.Println("x == y:", x == y)
  fmt.Println("x != y:", x != y)
  fmt.Println("x < y:", x < y)
  fmt.Println("x > y:", x > y)
  fmt.Println("x <= y:", x <= y)
  fmt.Println("x >= y:", x >= y)
}