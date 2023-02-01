package main

import "fmt"

func main() {
	ch := make(chan string, 1)
	sendChan(ch)
	receiveChan(ch)
}

func sendChan(ch chan<- string) { // 송신 채널 파라미터
	ch <- "Data" // ch 라는 채널에 "Data" 를 보낸다. 즉 ch 는 "송신 전용 채널"
	// x := <-ch // error
}

func receiveChan(ch <-chan string) { // 수신 채널 파라미터
	data := <-ch //  "수신 전용 채널"
	fmt.Println(data)
}
