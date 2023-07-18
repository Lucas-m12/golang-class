package main

func main() {
	index := 0
	for index <= 10 {
		println(index)
		index++
	}
	println("------------------------------")
	for index2 := 0; index2 <= 10; index2++ {
		println(index2)
	}
	println("------------------------------")
	names := [3]string{"Lucas", "Laura", "Lays"}
	for _, name := range names {
		println(name)
	}
	println("------------------------------")
	for _, letter := range "palavra" {
		println(string(letter))
	}
	println("------------------------------")
	user := map[string]string{
		"name":    "Lucas",
		"surname": "Marinho da Silva",
	}
	for _, value := range user {
		println(value)
	}
}
