package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) calArea() float64 {
	return s.side * s.side
}

func (c circle) calArea() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	calArea() float64
}

func info(x shape) {
	fmt.Println("My area : ", x.calArea())
}

func main() {
	s1 := square{5}
	c1 := circle{7}

	info(s1)
	info(c1)

}
