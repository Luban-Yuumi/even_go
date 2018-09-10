package main

import (
	"container/ring"
	"fmt"
)

func main() {

	ring1 := ring.New(3)

	for i := 1; i <= 3; i++ {
		ring1.Value = i
		ring1 = ring1.Next()
	}

	ring2 := ring.New(3)

	for i := 4; i <= 6; i++ {
		ring2.Value = i
		ring2 = ring2.Next()
	}

	r := ring1.Link(ring2)

	fmt.Printf("ring length = %d\n", r.Len())

	r.Do(func(p interface{}) {
		fmt.Print(p.(int))
		fmt.Print(",")
	})

	fmt.Println()

	fmt.Printf("current ring is %v\n", r.Value)

	fmt.Printf("next ring is %v\n", r.Next().Value)

	fmt.Printf("prev ring is %v\n", r.Prev().Value)

	// ring 的遍历
	for p := r.Next(); p != r; p = p.Next() {
		fmt.Print(p.Value.(int))
		fmt.Print(",")
	}

}
