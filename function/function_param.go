package main

import "fmt"

func main() {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1) //值类型未做出修改
	array3 := [3][]string{
		[]string{"as", "bs", "cs"},
		[]string{"as", "bs", "cs"},
		[]string{"as", "bs", "cs"},
	}
	array4 := modiftSlice(array3)
	fmt.Printf("The modified array: %v\n", array4)
	fmt.Printf("The original array: %v\n", array3) //值类型但元素是切片对切片做出修改会改变，但是改变数组不会改变
}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}
func modiftSlice(a [3][]string) [3][]string {
	a[1][1] = "as"
	return a
}
