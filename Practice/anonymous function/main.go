package main

import "fmt"

func main() {

	calc := func(a int, b int) int {
		return a + b
	}

	fmt.Println(calc(1, 2))
}
