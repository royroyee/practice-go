package main

import "fmt"

func main() {
	m := make(map[string]string) // 맵 생성

	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"

	m["이화랑"] = "청주시 상당구"

	fmt.Printf("송하나 주소 %s \n", m["송하나"])
}
