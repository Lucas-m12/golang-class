package main

import "fmt"

type User struct {
	name    string
	age     uint8
	address Address
}

type Address struct {
	city  string
	state string
}

func main() {
	var user User
	user.name = "Lucas"
	user.age = 22
	user2 := User{
		age:  14,
		name: "Laura",
		address: Address{
			city:  "Cajueiro",
			state: "Alagoas",
		},
	}
	fmt.Println(user, user2)

}
