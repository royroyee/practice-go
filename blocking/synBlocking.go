package main

import (
	"fmt"
	"time"
)

// Synchronous & Blocking

func synBlockingOperation() {
	fmt.Println("start synchronous blocking operation")
	time.Sleep(time.Second * 3)
	fmt.Println("end synchronous blocking operation!")
}

func main() {
	fmt.Println("start main function...")

	synBlockingOperation()

	fmt.Println("end main function!")
}
