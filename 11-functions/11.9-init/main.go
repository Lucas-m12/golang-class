package main

var n int

func init() {
	n = 10
	println("Init function being executed")
}

func main() {
	println("Main function being executed")
	println(n)
}
