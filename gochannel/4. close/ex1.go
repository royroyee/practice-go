package main

func main() {
	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다.
	// Point : 채널을 닫게 되면, 해당 채널로는 더이상 송신할 순 없지만, 수신은 계속 가능하다!
	close(ch)

	// 채널 수신
	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("더이상 데이터 없음")
	}

}
