package main

import "fmt"

func main() {
	var v1 int8 = 10
	var v2 *int8 = &v1

	fmt.Println(v1, *v2)
	v1++
	fmt.Println(v1, *v2)
}
