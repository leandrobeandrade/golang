package main

import (
	mypackage "digitalocean/packages" // pacote interno
	"fmt"                             // pacote padrão

	"github.com/spf13/cobra" // pacote externo
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")
			mypackage.PrintHello()
		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}
