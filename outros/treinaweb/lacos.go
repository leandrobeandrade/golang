package main

import "fmt"

func main() {
  idade1, idade2, idade3 := 20, 35, 10
  bol := true

  fmt.Println("### if/else ###")
  if idade1 >= 18 && true {
    fmt.Println("Maior! ", idade2)

    if bol {
      fmt.Println("Verdadeiro!")
    }
  } else {
    fmt.Println("Menor! ", idade3)
  }

  fmt.Println("\n### for ###")
  for i := 0; i < 10; i++ {
    fmt.Println("i => ", i)
  }
}