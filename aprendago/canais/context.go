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

	lang1 := lang{language: "Go"}
	lang2 := lang{
		language: "Python",
	}

	ctx2 := context.WithValue(context.Background(), lang1, "Go")
	withValue(ctx2, "language")

	ctx2 = context.WithValue(ctx2, lang2, "Python")
	withValue(ctx2, "language")
	withValue(ctx2, "other")
}

func withCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")
}

func withValue(ctx context.Context, s string) {
	if v := ctx.Value(s); v != nil {
		fmt.Println("found value:", v)
		return
	}
	fmt.Println("key not found:", s)
}

type lang struct {
	language string
}
