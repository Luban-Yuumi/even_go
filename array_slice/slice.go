package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Println(a)

	b := make([]int, 5, 8)
	fmt.Println(len(b))
	fmt.Println(cap(b))
	fmt.Println(b)

	c := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(len(c))
	fmt.Println(cap(c))
	fmt.Println(c)

	d := c[3:6]
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Println(d)

	d = d[0:cap(d)]
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Println(d)

	d = append(d, 9)
	fmt.Println(len(d))
	fmt.Println(cap(d))
	fmt.Println(d)

	d = append(d, 9)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	d = append(d, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1)
	fmt.Println(len(d))
	fmt.Println(cap(d))

	e := c[3:6]
	f := c[3:6]
	fmt.Println(e[1])
	fmt.Println(f[1])

	f = append(f, 1, 2, 3, 4, 5, 6)
	e[1] = 30
	fmt.Println(e[1])
	fmt.Println(f[1])
}
