package main

import (
	"fmt"
	"os"
)

func main() {
	c := make(chan int, 3)
	c <- 3
	c <- 2
	c <- 1
	close(c)
	for {
		if value, ok := <-c; ok {
			fmt.Println(value)
		} else {
			os.Exit(0)
		}
	}
}
