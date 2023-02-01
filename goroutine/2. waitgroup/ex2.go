package main

import (
	"fmt"
	"sync"
)

func printer(val int) {
	for i := 0; i < 3; i++ {
		fmt.Println(val)
	}
}

func main() {
	wg := sync.WaitGroup{} // WaitGroup 생성
	go func() {
		wg.Add(1)
		defer wg.Done()
		printer(1)
	}()
	go func() {
		wg.Add(1)
		defer wg.Done()
		printer(2)
	}()
	fmt.Println("Hello!")
	wg.Wait()
}
