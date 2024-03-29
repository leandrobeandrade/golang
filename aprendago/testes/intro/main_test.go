package main

import (
	"fmt"
	"testing"
)

func ExampleSoma() {
	fmt.Println(Soma(4, 2, 1)) // 7
	fmt.Println(Soma(4, 2, 2)) // 8
	fmt.Println(Soma(4, 2, 3)) // 9
}

func TestSoma(t *testing.T) {
	teste := Soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := Multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Expected:", resultado, "Got:", teste)
	}
}

// - struct test, fields: data []int, answer int
// - tests := []test{[]int{}, int}
// - range tests

type test struct {
	data   []int
	answer int
}

func TestSomaEmTabela(t *testing.T) {
	tests := []test{
		{data: []int{1, 2, 3}, answer: 6},
		{[]int{10, 11, 12}, 33},
		{[]int{-5, 0, 5, 10}, 10},
	}
	for _, v := range tests {
		z := Soma(v.data...)
		if z != v.answer {
			t.Error("Expected:", v.answer, "Got:", z)
		}
	}
}

func BenchmarkSoma(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Soma(1, 1)
	}
}

func BenchmarkMultiplica(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Multiplica(2, 1)
	}
}
