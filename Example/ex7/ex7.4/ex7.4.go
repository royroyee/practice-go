package main

import "fmt"

func Divide(a, b int) (int, bool) {

	if b == 0 {
		return 0, false
	}

	return a / b, true

}

func main() {

	c, success := Divide(5, 3)
	fmt.Println(c, success)

	c, false := Divide(1, 0)
	fmt.Println(c, false)
}
