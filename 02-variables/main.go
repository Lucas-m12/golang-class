package main

import "fmt"

func main() {
	var v1 string = "lucas"
	v2 := "marinho"
	var (
		v3 string = "da"
		v4 string = "silva"
	)
	fmt.Println(v1, v2, v3, v4)
	v5, v6 := "Laura", "Beatriz"
	const c1 string = "Marinho"
	fmt.Println(v5, v6, c1)
	v5, v6 = v6, v5
	fmt.Println(c1, v5, v6)
}
