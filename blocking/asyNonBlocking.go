package main

import (
	"fmt"
	"time"
)

func asyNonBlockingOperation(c chan string) {
	fmt.Println("start asynchronous non-blocking operation...")
	time.Sleep(time.Second * 3)
	c <- "end asynchronous non-blocking operation"
}

func main() {
	fmt.Println("start main function")

	c := make(chan string, 1)

	go asyNonBlockingOperation(c)

	for {
		select {
		case result := <-c:
			fmt.Println(result)
			return
		default:
			fmt.Println("continue main function...")
			time.Sleep(time.Second)
		}
	}
}
