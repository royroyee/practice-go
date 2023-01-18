package main

import "fmt"

type account struct {
	balance int
}

// 함수
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메서드
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {

	a := &account{100}

	withdrawFunc(a, 30) // 함수 형태 호출

	a.withdrawMethod(30)

	fmt.Printf("%d \n", a.balance)

}
