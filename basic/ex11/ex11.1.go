package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	// 기본 for문
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	// while문 처럼
	for i := 0; i < 10; {
		fmt.Println(i)
		i++
	}

	var k int = 0
	// 무한루프
	for {
		if k >= 10 {
			break
		}
		fmt.Println(k)
		k++
	}

	for {
		fmt.Print("숫자 하나를 입력해주세요. :  ")
		var number int
		_, err := fmt.Scanln(&number)

		if err != nil {
			fmt.Println("숫자를 입력하셔야 합니다.")
			stdin.ReadString('\n')
			continue
		}

		fmt.Printf("입력하신 숫자는 %d 입니다.", number)

		break
	}
	fmt.Println("for문 종료!")
}
