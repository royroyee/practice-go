package main

import "fmt"

var x int // main package 안에 있다면 모두 접근 가능한 변수

func main() {

	x := 7
	fmt.Printf("%T", x)

}
