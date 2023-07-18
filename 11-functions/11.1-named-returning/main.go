package main

func calculate(n1, n2 int) (sum, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	sum, sub := calculate(2, 2)
	println(sum, sub)
}
