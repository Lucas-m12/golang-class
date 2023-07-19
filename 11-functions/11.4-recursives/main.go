package main

import "time"

func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}
	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {
	var position uint = 12
	for index := uint(0); index < position; index++ {
		value := fibonacci(index)
		println(value)
		time.Sleep(time.Second)
	}
}
