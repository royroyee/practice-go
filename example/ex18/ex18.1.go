package main

import "fmt"

func main() {

	//	arr1 := [5]int{1, 2, 3, 4, 5}

	arr2 := []int{7, 8, 9, 0}

	arr2 = append(arr2, 1)

	for _, v := range arr2 {
		fmt.Println(v)
	}
}
