package main

import "fmt"

func main() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()
	// 위의 고루틴이 끝날 때까지 대기(true false 등과는 관계없음)
	<-done
}
