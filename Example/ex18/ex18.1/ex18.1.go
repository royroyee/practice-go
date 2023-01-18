package main

import "fmt"

func main() {

	// 배열 (정적)
	// 슬라이스 (동적) (원소 수를 입력하지 않는다. 배열에서 쓰던 [...] 도 허용 x)
	arr2 := []int{7, 8, 9, 0}

	// 슬라이스는 append 메서드 제공
	arr2 = append(arr2, 1, 2, 3) // 여러값을 한 번에 추가도 가능하다

	for _, v := range arr2 {
		fmt.Println(v)
	}
}
