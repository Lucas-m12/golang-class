package main

func closure() func() {
	text := "In closure function"
	function := func() {
		println(text)
	}
	return function
}

func main() {
	text := "In main function"
	println(text)
	newFunction := closure()
	newFunction()
}
