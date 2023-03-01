package main

import (
	"fmt"
	"time"
)

func blockingOperatinon() {
	fmt.Println("start blocking operation...")
	time.Sleep(time.Second * 3)
	fmt.Println("end blocking operation")
}

func nonBlockingOperation(c chan string) {
	fmt.Println("start non-blocking operation...")
	time.Sleep(time.Second * 3)
	c <- "end non-blocking operatinon!"
}

func main() {
	// blocking
	blockingOperatinon()

	// non-blocking
	fmt.Println("start non-blocking operation...(main)")
	c := make(chan string)
	go nonBlockingOperation(c)

	// non-blocking 작업 결과를 기다리지 않고 다른 작업 수행 가능 (여기선 nonBlockingOperation 함수에 따라 약 3초 동안 수행)
	fmt.Println("do other work while waiting for non-blocking operation")

	// non-blocking result
	fmt.Println(<-c)
}
