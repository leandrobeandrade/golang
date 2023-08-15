package main

import ("fmt")

func porValor(c int) {
  c = 20
}

func porReferencia(d *int) {
  *d = 20
}

func main() {
  a := 10
  var b *int = &a

  fmt.Println("\n### ponteiros ###")
  fmt.Println("A ender ", &a)
  fmt.Println("B ender ", b)
  fmt.Println("A valor ", *b)

  *b = 20

  fmt.Printf("\nNovo valor a %d\n", a)
  fmt.Printf("Novo valor b %d\n", *b)

  var c int = 10
  var d int = 10

  fmt.Println("\n### passagem por valor ###")
  fmt.Printf("Valor c %d\n", c)
  porValor(c)
  fmt.Printf("Valor c %d\n", c)

  fmt.Println("\n### passagem por referencia ###")
  fmt.Printf("Valor d %d\n", d)
  porReferencia(&d)
  fmt.Printf("Valor d %d\n", d)

  fmt.Println("\n============================================\n")

  i, j := 42, 2701

  p := &i         // endereço de memória para i (ponteiro de i)
  fmt.Println(*p) // lê o valor de i através do ponteiro
  *p = 21         // altera o valor de i através do ponteiro
  fmt.Println(i)  // lê o novo valor de i

  p = &j         // ponteiro de J
  *p = *p / 37   // divide o valor de J usando o ponteiro
  fmt.Println(j) // verifica o novo valor de J
}
