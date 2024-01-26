package main

import "fmt"

var arr1 [5]int                     // array de inteiros de tamanho 5
var arr2 = [6]int{1, 2, 3, 4, 5, 6} // array inicializado
var sli1 []int                      // slice de inteiros de tamanho indeterminado
var sli2 = []int{1, 2, 3}           // slice inicializado

func main() {
	fmt.Println("=== Array ===")
	arr1[0] = 1
	arr1[1] = 10
	fmt.Println(arr1[0], arr1[1])
	fmt.Println(arr1, arr2)
	fmt.Printf("%T %T\n", arr1, arr2)
	fmt.Println(len(arr1))

	fmt.Println("\n=== Slice ===")
	sli1 = append(sli1, 1, 2, 3)
	fmt.Println("Slice1", sli1)

	sli3 := append(sli2, 2)
	fmt.Println("Slice3 - Slice2", sli3, "\t", sli2)
	fmt.Printf("%T %T\n", sli3, sli2)
	fmt.Println(len(sli3))

	fmt.Println("\n=== Slice com range ===")
	sli4 := []string{"banana", "maçã", "jaca"}
	for idx, val := range sli4 {
		fmt.Println("No índice:", idx, "temos o valor:", val)
	}
	fmt.Println()

	sli5 := []int{1, 2, 3}
	var total int = 1
	for _, val := range sli5 {
		total = total * val
	}
	fmt.Println("Valor total:", total)

	fmt.Println("\n=== Obtendo valores ===")
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
	fmt.Println(sabores[0])
	fmt.Println(sabores[0:2])
	fmt.Println(sabores[:2])
	fmt.Println(sabores[2:])
	// fmt.Println(sabores[2:len(sabores)])
	fmt.Println(sabores[:])

	fmt.Println("\n=== Adicionando valores ===")
	nums1 := []int{1, 2, 3}
	nums2 := []int{8, 9, 10}

	nums1 = append(nums1, 5, 6, 7)
	fmt.Println(nums1)
	nums1 = append(nums1, nums2...)
	fmt.Println(nums1)

	fmt.Println("\n=== Deletando valores ===")
	sabores = append(sabores[:2], sabores[4:]...) // deleta valores (abacaxi, quatroqueijos)
	fmt.Println(sabores[:])

	fmt.Println("\n=== Método make() ===")
	sli6 := make([]int, 5, 10)
	fmt.Println(sli6, len(sli6), cap(sli6))
	sli6[0], sli6[1], sli6[2], sli6[3] = 1, 2, 3, 4
	fmt.Println(sli6, len(sli6), cap(sli6))
	// sli6[5] = 25 ERRO - o len (tamanho) é de 5, ou seja, vai até sli6[4]
	sli6 = append(sli5, 4) // append aumenta o len do slice
	fmt.Println(sli6, len(sli6), cap(sli6))
	sli6 = append(sli6, 5)
	sli6 = append(sli6, 6)
	sli6 = append(sli6, 7)
	sli6 = append(sli6, 8)
	sli6 = append(sli6, 9)
	sli6 = append(sli6, 10)
	sli6 = append(sli6, 11)
	fmt.Println(sli6, len(sli6), cap(sli6)) // aumenta a capacidade do slice se ultrapassada

	fmt.Println("\n=== Slices múltiplos ===")
	ss := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(ss)
	fmt.Println(ss[1])
	fmt.Println(ss[1][1])
}
