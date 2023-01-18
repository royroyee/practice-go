package main

import "fmt"

func Divide(a int, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}

	result = a / b
	success = true
	return

}

func main() {

	a, b := Divide(3, 4)
	fmt.Println(a, b)
}
