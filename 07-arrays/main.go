package main

import "fmt"

func main() {
	var array1 [4]int
	array1[0] = 2
	array1[1] = 2
	array1[2] = 2
	array1[3] = 2
	fmt.Println(array1)
	array2 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(array2)

	slice := []int{1, 2}
	sliceAppend := append(slice, 3, 4)
	fmt.Println(slice, sliceAppend)
}
