package main

import (
	"fmt"
)

var test = map[interface{}]int{
	"a":        1,
	"[]int{2}": 2,
	3:          3,
}

func main() {
	fmt.Printf("%v", test[[]int{2}]) //在编译时不会出错，并且直接fmt.Printf("%v", test)也不会出错 ，但是一旦要取对应的值就会报错
}
