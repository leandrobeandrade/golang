package generics

import "golang.org/x/exp/constraints"

func Soma1[T int64 | float64](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

type Varios interface {
	int64 | float64 | rune
}

// Soma2() recebe uma interface (CONSTRAINTS) com vários tipos suportados
func Soma2[T Varios](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

type Number interface {
	~float64 | ~rune
}

// NumInt implementa outra interface (Number)
type NumInt int

func Soma3[T NumInt](m map[string]T) T {
	var soma T

	for _, v := range m {
		soma += v
	}
	return soma
}

// Soma4() -> ou (number1 T, number2 T)
func Soma4[T Number](number1, number2 T) T {
	return number1 + number2
}

func Compara[T comparable](number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}

func Max[T constraints.Ordered](number1 T, number2 T) T {
	if number1 > number2 {
		return number1
	}
	return number2
}

type Stringer interface {
	String() string
}

type MyString string

func (m MyString) String() string {
	return string(m)
}

func Concat[T Stringer](vals []T) string {
	resultado := ""
	for _, val := range vals {
		resultado += val.String()
	}
	return resultado
}
