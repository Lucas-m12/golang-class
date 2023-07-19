package main

func example1() {
	println("Executing the example 1")
}

func example2() {
	println("Executing the example 2")
}

func isApprovedStudent(n1, n2 float32) bool {
	defer println("calculated media. Result to be return")
	println("verifing if student is approved")
	if (n1+n2)/2 >= 6 {
		return true
	} else {
		return false
	}
}

func main() {
	result := isApprovedStudent(5.5, 8.7)
	println(result)
}
