package main

import "fmt"

//func main() {
//	c := make(chan int)
//	c <- 1           // 수신 루틴이 없으므로 데드락
//	fmt.Println(<-c) // 코멘트해도 데드락(별도의 고루틴이 없기 때문이다.)
//
//}

func main() {
	ch := make(chan int, 1) // 버퍼 채널 사용

	// 수신자가 없더라도 보낼 수 있음
	ch <- 101

	fmt.Println(<-ch)
}
