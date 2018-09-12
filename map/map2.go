package main

import (
	"fmt"
)

var test map[int]int

func main() {
	fmt.Println(test) //只声明了的map可以读
	test['a'] = 1     //这个举动会抛panic,assignment to entry in nil map
}
