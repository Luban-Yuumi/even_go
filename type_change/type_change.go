package main

import (
	"fmt"
)

func main() {
	//int类型转int类型
	var a int16 = -255
	b := int8(a)
	//int类型转string类型
	c := 1
	d := string(c)
	//string类型转切片类型
	e := "你好呀"
	f := []byte(e)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println(f)
}
