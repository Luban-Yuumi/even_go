package main

import (
	"errors"
	"fmt"
)

type operator func(a int, b int) int
type calcutar func(x int, y int) int

func main() {
	x := 20
	y := 20
	err, op := getCalcutar(operator1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(op(x, y))
}

func getCalcutar(op operator) (error, calcutar) {
	if op == nil {
		return errors.New("there is on operator"), nil
	}
	return nil, func(x int, y int) int {
		return op(x, y)
	}
}

func operator1(x int, y int) int {
	return x + y
}
