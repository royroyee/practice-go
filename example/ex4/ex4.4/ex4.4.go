package main

import "fmt"

func main() {

	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // float64 -> int
	d := float64(a * c) // int -> float

	var e int64 = 7
	f := int64(d) * e // float64 -> int64

	var g int = int(b * 3) // float64 -> int
	var h int = int(b) * 3 // float64 -> int g랑 값 다름!
	fmt.Println(g, h, f)
}
