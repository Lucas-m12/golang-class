package main

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func write(text string, numbers ...int) {
	println(text, sum(numbers...))
}

func main() {
	var numbers []int = []int{2, 4, 6, 8, 10, 12}
	result := sum(numbers...)
	println(result)
	write("resultado da soma:", numbers...)
}
