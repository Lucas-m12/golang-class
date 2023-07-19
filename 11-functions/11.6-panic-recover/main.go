package main

func recoveryExecution() {
	if r := recover(); r != nil {
		println("Execution recovery with success!")
	}
}

func isApprovedStudent(n1, n2 float32) bool {
	defer recoveryExecution()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("The media is 6!")
}

func main() {
	result := isApprovedStudent(6, 6)
	println(result)
}
