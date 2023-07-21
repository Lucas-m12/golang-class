package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radio float64
}

type Shape interface {
	area() float64
}

func (rectangle Rectangle) area() float64 {
	return rectangle.height * rectangle.width
}

func (circle Circle) area() float64 {
	return math.Pi * math.Pow(circle.radio, 2)
}

func writeArea(shape Shape) {
	fmt.Printf("A area da forma é %0.2f \n", shape.area())
}

func main() {
	rectangle := Rectangle{height: 10, width: 15}
	writeArea(rectangle)
	circle := Circle{radio: 10}
	writeArea(circle)
}
