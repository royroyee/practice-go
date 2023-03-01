package main

import (
	"fmt"
	"time"
)

func synNonBlockingOperation(c chan string) {
	fmt.Println("start synchronous non-blocking operation...")
	time.Sleep(time.Second * 3)
	c <- "end synchronous non-blocking operation!"
}

func main() {
	fmt.Println("start main function")

	c := make(chan string)

	go synNonBlockingOperation(c)

	fmt.Println("other work")

	fmt.Println(<-c)

	fmt.Println("end main function!")

}
