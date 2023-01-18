package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func printArea(x shape) {
	fmt.Println("My area is : ", x.area())
}

func main() {
	s1 := square{10}
	s2 := circle{5}
	printArea(s1)
	printArea(s2)
}
