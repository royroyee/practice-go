package main

import (
	"fmt"
	"time"
)

func asynchronousBlockingOperation(c chan string) {
	fmt.Println("start asynchronous blocking operation...")
	time.Sleep(time.Second * 3)
	c <- "end asynchronous blocking operation!"
}

func main() {
	fmt.Println("start main function...")

	c := make(chan string)

	asynchronousBlockingOperation(c)

	fmt.Println(<-c)

	fmt.Println("end main function!")
}
