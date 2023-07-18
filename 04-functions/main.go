package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalc(n1, n2 int8) (int8, int8) {
	resultSum := sum(n1, n2)
	resultMult := n1 * n2
	return resultSum, resultMult
}

func main() {
	result := sum(127, 127)
	fmt.Println(result)

	var function = func(text string) string {
		fmt.Println(text)
		return text
	}
	function("Hello World!")

	resultSum, resultMult := mathCalc(10, 10)
	fmt.Println(resultSum, resultMult)
}
