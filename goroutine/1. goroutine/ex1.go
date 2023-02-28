package main

import (
	"fmt"
	"time"
)

func PrintAlphabet() {
	alphabet := "abcdefghijklmnop"

	for _, v := range alphabet {
		time.Sleep(200 * time.Millisecond) // nothing in particular
		fmt.Printf("%c", v)
	}
}

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(200 * time.Millisecond) // nothing in particular
		fmt.Printf("%d", i)
	}
}

func main() {
	PrintAlphabet()
	fmt.Println("")

	PrintNumbers()
	fmt.Println("")

	go PrintAlphabet()
	go PrintNumbers()
	time.Sleep(3 * time.Second) // If not, functions executed with go keyword are terminated without being executed.
	// This is because the main routine does not wait for the result of the subroutine.
}
