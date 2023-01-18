package main

import "fmt"

const X int = 3

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	arr2 := [X]int{1, 2, 3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])

	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr2[i])

	}

	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])

	}

	for i, v := range arr3 {
		fmt.Println("인덱스 : ", i, "값 : ", v)
	}
}
