package main

import (
	"errors"
	"fmt"
)

func main() {

	// INTEGER NUMBERS
	var n1 int8 = 120
	fmt.Println(n1)

	var n2 uint32 = 1000
	fmt.Println(n2)

	var n3 rune = 1000
	fmt.Println(n3)

	// FLOAT NUMBERS
	var n4 float32 = 12.77
	var n5 float64 = 12000.77
	fmt.Println(n4, n5)

	// STRINGS
	var s1 string = "Lucas Marinho da Silva"
	s2 := "Laura Beatriz Marinho da Silva"
	fmt.Println(s1, s2)
	char := 'L'
	var char2 int8 = 'U'
	fmt.Println(char, char2)

	// BOOLEAN
	var b1 bool = true
	var b2 bool = false
	fmt.Println(b1, b2)

	// ERROR
	var err error = errors.New("error test")
	fmt.Println(err)

}
