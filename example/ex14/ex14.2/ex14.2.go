package main

import "fmt"

func main() {

	var a int = 10
	b := 10

	var p1 *int = &a
	var p2 *int = &b

	fmt.Printf("a == b : %v\n", a == b)
	fmt.Printf("p1 == p2 : %t\n", p1 == p2)
}
