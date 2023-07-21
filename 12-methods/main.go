package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func (user User) save() {
	fmt.Printf("Saving user data from %s on database", user.name)
}

func main() {
	user1 := User{age: 23, name: "Lucas Marinho da Silva"}
	user1.save()
}
