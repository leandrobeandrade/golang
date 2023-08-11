package main

import "fmt"

type Person struct {
    name string
    age int
    email string
}

func porValor(c int) {
    c = 20
}

func porReferencia(d *int) {
    *d = 20
}

func main() {
    age := 10
    fmt.Println("Olá Mundo GO >>>>", age)

    fmt.Println("\n### if/else ###")
    if age >= 18 {
        fmt.Println("Maior!")
    } else {
        fmt.Println("Menor!")
    }

    fmt.Println("\n### for ###")
    for i := 0; i < 10; i++ {
        fmt.Println("i => ", i)
    }

    person := Person {
        name: "Fulano",
        age: 35,
        email: "fulano@gmail.com",
    }

    fmt.Println("\n### struct ###")
    fmt.Println("A pessoa tem o nome ", person.name, " a idade ", person.age, " anos e o email: ", person.email)

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
}