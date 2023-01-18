package main

import (
	"fmt"
)

func main() {
	// slice
	// conmpoiste literal - slice literal

	x := []int{7, 9, 31} // 가장 기본적인 슬라이스 선언 방법
	fmt.Println(x)

	for i, _ := range x { // 반환 값 : index , value
		fmt.Println(i, "-", x[i])
	}

	y := make([]int, 0, 100) // make를 이용한 슬라이스 선언 방법 len =0, cap=100
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 7)
	y = append(y, 31)

	for i, v := range y {
		fmt.Println(i, "-", v)
	}

	fmt.Println(y)

}
