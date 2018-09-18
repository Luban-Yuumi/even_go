package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	test := make(<-chan int, 1)
	time.AfterFunc(time.Second, func() {
		close(test)
	})
	for {
		select {
		case _, ok := <-test: //程序是不会进入1的，因为test没有关闭之前 case 处于阻塞状态，所以会执行default代码块，关闭后ok才等于false
			if ok {
				fmt.Println("1")
			} else {
				fmt.Println("2")
				os.Exit(0)
			}
		default:
			fmt.Println("3")
		}
	}
}
