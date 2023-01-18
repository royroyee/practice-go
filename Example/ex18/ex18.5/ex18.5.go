package main

import "fmt"

func main() {

	slice1 := make([]int, 3, 5)

	slice2 := make([]int, 5)

	slice1 = append(slice1, 1)
	slice2 = append(slice2, 1)

	fmt.Println("slice1[3] : ", slice1[3])

	for j, k := range slice1 {
		fmt.Println(j, k)
	}

	for i, v := range slice2 {
		fmt.Println(i, v)
	}
}
