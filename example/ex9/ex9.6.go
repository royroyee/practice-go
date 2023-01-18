package main

import "fmt"

func getMyAge() (int, bool) {
	return 10, true
}

func main() {

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("You are young", age)
	}

	// fmt.Println(age) 초기문은 if문 안에서만 접근이 가능하다.
}
