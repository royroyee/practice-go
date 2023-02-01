package main

func main() {
	// 정수형 채널 생성
	ch := make(chan int)

	go func() {
		ch <- 123 // ch 라는 채널에 123을 보낸다.
	}()

	var i int
	i = <-ch   // i는 ch 라는 채널로부터 123을 받는다.
	println(i) // 123
}
