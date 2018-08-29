package main

import (
	"learnGo/learn_internal/lib"
	// "learnGo/learn_internal/lib/internal" 会报错 demo.go:5:2: use of internal package not allowed
)

func main() {
	demo1.Hello("nshu")
	// internal.Hello("nshu")
}
