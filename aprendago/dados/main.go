package main

import "fmt"

var arr1 [5]int                     // array de inteiros de tamanho 5
var arr2 = [6]int{1, 2, 3, 4, 5, 6} // array inicializado
var sli1 []int                      // slice de tamanho indeterminado
var sli2 = []int{1, 2, 3}           // slice inicializado

func main() {
	fmt.Println("=== Array ===")
	arr1[0] = 1
	arr1[1] = 10
	fmt.Println(arr1[0], arr1[1])
	fmt.Println(arr1, arr2)
	fmt.Printf("%T %T\n", arr1, arr2)
	fmt.Println(len(arr1))

	fmt.Println("=== Slice ===")
	sli1 = append(sli1, 1, 2, 3)
	fmt.Println("Slice1", sli1)

	sli3 := append(sli2, 2)
	fmt.Println("Slice3 - Slice2", sli3, "\t", sli2)
	fmt.Printf("%T %T\n", sli3, sli2)
	fmt.Println(len(sli3))

	fmt.Println("=== Slice com range ===")
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
	fmt.Println()

	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marguerita"}
	fmt.Println(sabores[0])
	fmt.Println(sabores[0:2])
	fmt.Println(sabores[:2])
	fmt.Println(sabores[2:])
	// fmt.Println(sabores[2:len(sabores)])
	fmt.Println(sabores[:])

	fmt.Println("\nAdiciona valores")
	nums1 := []int{1, 2, 3}
	nums2 := []int{8, 9, 10}

	nums1 = append(nums1, 5, 6, 7)
	fmt.Println(nums1)
	nums1 = append(nums1, nums2...)
	fmt.Println(nums1)

	fmt.Println("\nDeletando valores") // deleta valores (abacaxi, quatroqueijos)
	sabores = append(sabores[:2], sabores[4:]...)
	fmt.Println(sabores[:])
}
