package main

import "fmt"

func main() {

	arr1 := [5]int{1, 2, 3, 4, 5}
	arr2 := [5]int{100, 200, 300, 400, 500}

	for i, v := range arr1 {
		fmt.Printf("arr1[%d] : %d\n", i, v)
	}

	for i, v := range arr2 {
		fmt.Printf("arr2[%d] : %d\n", i, v)
	}

	arr1 = arr2

	for i, v := range arr1 {
		fmt.Printf("arr1[%d] : %d\n", i, v)
	}

	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr3 := range a {
		for _, v := range arr3 {
			fmt.Println(v, " ")
		}
	}
}
