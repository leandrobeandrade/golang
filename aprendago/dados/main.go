package main

import "fmt"

var arr1 [5]int                     // array de inteiros de tamanho 5
var arr2 = [6]int{1, 2, 3, 4, 5, 6} // array inicializado
// var sli1 []int                      // slice de tamanho indeterminado
var sli2 = []int{1, 2, 3} // array inicializado

func main() {
	fmt.Println("=== Array ===")
	arr1[0] = 1
	arr1[1] = 10
	fmt.Println(arr1[0], arr1[1])
	fmt.Println(arr1, arr2)
	fmt.Printf("%T %T\n", arr1, arr2)
	fmt.Println(len(arr1))

	fmt.Println("=== Slice ===")
	sli1 := append(sli2, 2)
	fmt.Println(sli1, sli2)
	fmt.Printf("%T %T\n", sli1, sli2)
	fmt.Println(len(sli1))
}
