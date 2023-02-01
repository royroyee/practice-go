package main

import (
	"fmt"
	"sync"
	"time"
)

func PrintAlphabet() {
	alphabet := "abcdefghijklmnop"

	for _, v := range alphabet {
		time.Sleep(200 * time.Millisecond) // nothing in particular
		fmt.Printf("%c", v)
	}

	wg.Done()
}

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond) // nothing in particular
		fmt.Printf("%d", i)
	}

	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2) // add two sub goroutine

	go PrintAlphabet()
	go PrintNumbers()

	wg.Wait() // wait until all sub routines are finished
}
