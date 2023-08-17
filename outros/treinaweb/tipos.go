package main

import "fmt"

func main() {
  fmt.Println("TIPOS INTEIRO")

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

  fmt.Println("\nTIPOS FLOAT")

  fmt.Println("\nTIPOS STRING")

  fmt.Println("\nTIPOS RUNE")
  // armazena códigos que representam caracteres Unicode
  var rune_  rune = 'A'
  fmt.Println("rune: ", rune_)

  // declarando e inicializando com um caractere Unicode
  emoji := "😀"
  fmt.Println("rune: ", emoji)

  fmt.Println("\nTIPOS BOOLEAN")
  bol := true
  fmt.Println(bol != true)
}