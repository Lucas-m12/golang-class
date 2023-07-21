package main

func invertSign(n int) int {
	return n * -1
}

func invertSignWithPointer(n *int) {
	*n = *n * -1
}

func main() {
	n := 10
	invertedNumber := invertSign(n)
	println(invertedNumber)
	println(n)

	println("-----------------------------")

	newIvertedNumber := -20
	println(newIvertedNumber)
	invertSignWithPointer(&newIvertedNumber)
	println(newIvertedNumber)

}
