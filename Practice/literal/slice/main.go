package main

import (
	"fmt"
)

func main() {
	// slice
	// conmpoiste literal - slice literal

	x := []int{7, 9, 31} // 가장 기본적인 슬라이스 선언 방법
	fmt.Println(x)

	y := make([]int, 0, 100) // make를 이용한 슬라이스 선언 방법 len =0, cap=100
	y = append(y, 8)
	y = append(y, 12)
	y = append(y, 7)
	y = append(y, 31)

	fmt.Println(y)

}
