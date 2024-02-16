package main

import (
	"context"
	"fmt"
)

func context_() {
	ctx := context.Background()

	fmt.Println("\ncontext:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	withCancel()

	l := lang("language")
	ctx2 := context.WithValue(context.Background(), l, "Go")
	withValue(ctx2, l)

	ctx2 = context.WithValue(ctx2, l, "Python")
	withValue(ctx2, l)
	withValue(ctx2, "Others")
}

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")
}

type lang string

func withValue(ctx context.Context, s lang) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}
