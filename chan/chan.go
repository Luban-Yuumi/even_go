package main

import (
	"fmt"
	"os"
)

func main() {
	a := 3
	c := make(chan int, 3)
	c <- a
	c <- 2
	c <- 1
	close(c) //如果不关闭通道，会引发daedlock
	for {
		if value, ok := <-c; ok {
			if value == 3 {
				value = 4
			}
			fmt.Println(value)
			fmt.Println(a)
		} else {
			os.Exit(0)
		}
	}
}
