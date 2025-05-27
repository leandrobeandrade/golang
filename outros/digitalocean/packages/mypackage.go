package __

import "fmt"

// função com a primeira letra maiúscula exporta a mesma (deixa global)
func PrintHello() {
	fmt.Println("Hello, Modules! This is mypackage speaking!")
}
